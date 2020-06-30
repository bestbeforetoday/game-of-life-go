package life

// Game of life. Constructor functions should be used to create instances.
type Game struct {
	cells     Cells
	neighbors func(Point) []Point
}

// State of a cell
type State bool

const (
	// Alive indicates that a cell is alive
	Alive State = true
	// Dead indicates that a cell is dead
	Dead State = false
)

// Cells represented by a map keyed by position and with a value of true if that position is a live cell
type Cells map[Point]State

// NewUnboundedGame creates an unbounded Game populated with the supplied live cells
func NewUnboundedGame(liveCellLocations []Point) Game {
	return Game{
		cells:     newCells(liveCellLocations),
		neighbors: unboundedNeighbors,
	}
}

// NewBoundedGame creates a bounded Game populated with the supplied live cells. The game is a rectangular area with
// corners at the minimum and maximum points. No cells can live outside of the game area.
func NewBoundedGame(liveCellLocations []Point, min Point, max Point) Game {
	return Game{
		cells:     newCells(liveCellLocations),
		neighbors: boundedNeighbors(min, max)}
}

func newCells(locations []Point) Cells {
	cells := make(Cells, len(locations))

	for _, location := range locations {
		cells[location] = Alive
	}

	return cells
}

// Cells in the game
func (g Game) Cells() Cells {
	return g.cells
}

// Next iteration of the game
func (g Game) Next() Game {
	nextCells := make([]Point, 0, len(g.cells))

	g.forEachLiveCell(func(location Point) {
		if g.isSurvivor(location) {
			nextCells = append(nextCells, location)
		}
		nextCells = append(nextCells, g.neighbourBirths(location)...)
	})

	return NewUnboundedGame(nextCells)
}

func (g Game) forEachLiveCell(fn func(Point)) {
	for cell := range g.cells {
		fn(cell)
	}
}

func (g Game) isSurvivor(location Point) bool {
	liveNeighbourCount := g.liveNeighbourCount(location)
	return liveNeighbourCount == 2 || liveNeighbourCount == 3
}

func (g Game) liveNeighbourCount(location Point) (count int) {
	g.forEachNeighbour(location, func(neighbour Point) {
		if g.cells[neighbour] == Alive {
			count++
		}
	})
	return count
}

func (g Game) forEachNeighbour(location Point, fn func(Point)) {
	for _, neighbour := range g.neighbors(location) {
		fn(neighbour)
	}
}

func (g Game) neighbourBirths(location Point) (locations []Point) {
	g.forEachNeighbour(location, func(neighbour Point) {
		if g.isBorn(neighbour) {
			locations = append(locations, neighbour)
		}
	})
	return locations
}

func (g Game) isBorn(location Point) bool {
	return g.cells[location] == Dead && g.liveNeighbourCount(location) == 3
}

func boundedNeighbors(min Point, max Point) func(Point) []Point {
	return func(p Point) []Point {
		possibleNeighbors := unboundedNeighbors(p)

		neighbors := make([]Point, 0, len(possibleNeighbors))
		for _, neighbour := range possibleNeighbors {
			if !neighbour.LessThan(min) && !neighbour.GreaterThan(max) {
				neighbors = append(neighbors, neighbour)
			}
		}

		return neighbors
	}
}

func unboundedNeighbors(p Point) []Point {
	return []Point{
		{p.X - 1, p.Y - 1},
		{p.X - 1, p.Y},
		{p.X - 1, p.Y + 1},
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
		{p.X + 1, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X + 1, p.Y + 1},
	}
}
