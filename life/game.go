package life

// Game of life. Constructor functions should be used to create instances.
type Game struct {
	cells Cells
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

// NewGame creates a Game populated with the supplied live cells
func NewGame(liveCells []Point) Game {
	cells := make(Cells, len(liveCells))
	for _, location := range liveCells {
		cells[location] = Alive
	}

	return Game{cells}
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

	return NewGame(nextCells)
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
	forEachNeighbour(location, func(neighbour Point) {
		if g.cells[neighbour] == Alive {
			count++
		}
	})
	return count
}

func forEachNeighbour(location Point, fn func(Point)) {
	for _, neighbour := range neighbours(location) {
		fn(neighbour)
	}
}

func (g Game) neighbourBirths(location Point) (locations []Point) {
	forEachNeighbour(location, func(neighbour Point) {
		if g.isBorn(neighbour) {
			locations = append(locations, neighbour)
		}
	})
	return locations
}

func (g Game) isBorn(location Point) bool {
	return g.cells[location] == Dead && g.liveNeighbourCount(location) == 3
}

func neighbours(p Point) []Point {
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
