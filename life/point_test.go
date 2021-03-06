package life

import "testing"

func TestLessThan(t *testing.T) {
	t.Run("Lower X value is less", func(t *testing.T) {
		p := Point{-1, 0}
		if !p.LessThan(Point{0, 0}) {
			t.Error()
		}
	})

	t.Run("Lower Y value is less", func(t *testing.T) {
		p := Point{0, -1}
		if !p.LessThan(Point{0, 0}) {
			t.Error()
		}
	})

	t.Run("Identical values are not less", func(t *testing.T) {
		p := Point{0, 0}
		if p.LessThan(Point{0, 0}) {
			t.Error()
		}
	})
}

func TestGreaterThan(t *testing.T) {
	t.Run("Higher X value is greater", func(t *testing.T) {
		p := Point{1, 0}
		if !p.GreaterThan(Point{0, 0}) {
			t.Error()
		}
	})

	t.Run("Higher Y value is greater", func(t *testing.T) {
		p := Point{0, 1}
		if !p.GreaterThan(Point{0, 0}) {
			t.Error()
		}
	})

	t.Run("Identical values are not greater", func(t *testing.T) {
		p := Point{0, 0}
		if p.GreaterThan(Point{0, 0}) {
			t.Error()
		}
	})
}
