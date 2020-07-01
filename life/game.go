package life

// Game of life. Constructor functions should be used to create instances.
type Game struct {
	cells     Cells
	neighbors func(Point) []Point
}

// State of a cell
type State int

const (
	// Dead indicates that a cell is dead
	Dead State = iota
	// Alive indicates that a cell is alive
	Alive
)

// Cells represented by a map keyed by position and with a value of true if that position is a live cell
type Cells map[Point]State

// NewUnboundedGame creates an unbounded Game populated with the supplied live cells
func NewUnboundedGame(liveCellLocations []Point) *Game {
	return &Game{
		cells:     newCells(liveCellLocations),
		neighbors: unboundedNeighbors,
	}
}

// NewBoundedGame creates a bounded Game populated with the supplied live cells. The game is a rectangular area with
// corners at the minimum and maximum points. No cells can live outside of the game area.
func NewBoundedGame(liveCellLocations []Point, min Point, max Point) *Game {
	return &Game{
		cells:     newCells(liveCellLocations),
		neighbors: boundedNeighbors(min, max),
	}
}

func newCells(locations []Point) Cells {
	cells := make(Cells, len(locations))

	for _, location := range locations {
		cells[location] = Alive
	}

	return cells
}

// Cells in the game
func (g *Game) Cells() Cells {
	return g.cells
}

// Next iteration of the game
func (g *Game) Next() *Game {
	nextCells := make(Cells)

	g.forEachLiveCell(func(location Point) {
		if g.isSurvivor(location) {
			nextCells[location] = Alive
		}
		g.addNeighborBirths(nextCells, location)
	})

	return &Game{
		cells:     nextCells,
		neighbors: g.neighbors,
	}
}

func (g *Game) forEachLiveCell(fn func(Point)) {
	for cell := range g.cells {
		fn(cell)
	}
}

func (g *Game) isSurvivor(location Point) bool {
	liveNeighborCount := g.liveNeighborCount(location)
	return liveNeighborCount == 2 || liveNeighborCount == 3
}

func (g *Game) liveNeighborCount(location Point) (count int) {
	g.forEachNeighbor(location, func(neighbor Point) {
		if g.cells[neighbor] == Alive {
			count++
		}
	})
	return count
}

func (g *Game) forEachNeighbor(location Point, fn func(Point)) {
	for _, neighbor := range g.neighbors(location) {
		fn(neighbor)
	}
}

func (g *Game) addNeighborBirths(cells Cells, location Point) {
	g.forEachNeighbor(location, func(neighbor Point) {
		if cells[neighbor] != Alive && g.isBorn(neighbor) {
			cells[neighbor] = Alive
		}
	})
}

func (g *Game) isBorn(location Point) bool {
	return g.cells[location] == Dead && g.liveNeighborCount(location) == 3
}

func boundedNeighbors(min Point, max Point) func(Point) []Point {
	return func(p Point) []Point {
		possibleNeighbors := unboundedNeighbors(p)

		neighbors := make([]Point, 0, len(possibleNeighbors))
		for _, neighbor := range possibleNeighbors {
			if !neighbor.LessThan(min) && !neighbor.GreaterThan(max) {
				neighbors = append(neighbors, neighbor)
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
