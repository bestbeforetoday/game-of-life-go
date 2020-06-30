package life

import "testing"

func TestLessThan(t *testing.T) {
	t.Run("Lower X value is less", func(t *testing.T) {
		isLess := Point{-1, 0}.LessThan(Point{0, 0})
		if !isLess {
			t.Error()
		}
	})

	t.Run("Lower Y value is less", func(t *testing.T) {
		isLess := Point{0, -1}.LessThan(Point{0, 0})
		if !isLess {
			t.Error()
		}
	})

	t.Run("Identical values are not less", func(t *testing.T) {
		isLess := Point{0, 0}.LessThan(Point{0, 0})
		if isLess {
			t.Error()
		}
	})
}

func TestGreaterThan(t *testing.T) {
	t.Run("Higher X value is greater", func(t *testing.T) {
		isGreater := Point{1, 0}.GreaterThan(Point{0, 0})
		if !isGreater {
			t.Error()
		}
	})

	t.Run("Higher Y value is greater", func(t *testing.T) {
		isGreater := Point{0, 1}.GreaterThan(Point{0, 0})
		if !isGreater {
			t.Error()
		}
	})

	t.Run("Identical values are not greater", func(t *testing.T) {
		isGreater := Point{0, 0}.GreaterThan(Point{0, 0})
		if isGreater {
			t.Error()
		}
	})
}
