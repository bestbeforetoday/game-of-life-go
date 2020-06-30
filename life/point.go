package life

// Point represents X and X coordinates
type Point struct {
	X int
	Y int
}

// LessThan returns true if either the X or Y coordinates of this point are less than the corresponding coordinates of
// the supplied argument.
func (p Point) LessThan(other Point) bool {
	return p.X < other.X || p.Y < other.Y
}

// GreaterThan returns true if either the X or Y coordinates of this point are greater than the corresponding
// coordinates of the supplied argument.
func (p Point) GreaterThan(other Point) bool {
	return p.X > other.X || p.Y > other.Y
}
