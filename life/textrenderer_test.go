package life

import "testing"

func newRenderer(min Point, max Point) *TextRenderer {
	return &TextRenderer{
		Alive: '*',
		Dead:  ' ',
		Min:   min,
		Max:   max,
	}
}

func assertEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func TestSingleLiveCell(t *testing.T) {
	renderer := newRenderer(Point{0, 0}, Point{0, 0})
	game := NewUnboundedGame([]Point{{0, 0}})
	actual := renderer.Render(game)
	assertEquals(t, actual, "*")
}

func TestSingleDeadCell(t *testing.T) {
	renderer := newRenderer(Point{0, 0}, Point{0, 0})
	game := NewUnboundedGame(make([]Point, 0))
	actual := renderer.Render(game)
	assertEquals(t, actual, " ")
}

func TestRenderSingleRow(t *testing.T) {
	renderer := newRenderer(Point{0, 0}, Point{4, 0})
	game := NewUnboundedGame([]Point{
		{0, 0},
		{2, 0},
		{4, 0},
	})
	actual := renderer.Render(game)
	assertEquals(t, actual, "* * *")
}

func TestRenderMultipleRows(t *testing.T) {
	renderer := newRenderer(Point{0, 0}, Point{1, 2})
	game := NewUnboundedGame([]Point{
		{0, 0},
		{1, 2},
	})
	actual := renderer.Render(game)
	assertEquals(t, actual, " *\n  \n* ")
}

func TestRenderWindowOnLargerGame(t *testing.T) {
	renderer := newRenderer(Point{1, 1}, Point{2, 2})
	game := NewUnboundedGame([]Point{
		{2, 0},
		{1, 1},
		{1, 2},
		{2, 3},
	})
	actual := renderer.Render(game)
	assertEquals(t, actual, "* \n* ")
}
