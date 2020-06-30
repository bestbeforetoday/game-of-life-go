package life

import "strings"

// TextRenderer displays games as text
type TextRenderer struct {
	Alive rune
	Dead  rune
	Min   Point
	Max   Point
}

// Render a game as text
func (r TextRenderer) Render(game Game) string {
	builder := r.newRenderBuilder()
	cells := game.Cells()

	for y := r.Min.Y; y <= r.Max.Y; y++ {
		if y > r.Min.Y {
			builder.WriteRune('\n')
		}
		r.writeRow(builder, cells, y)
	}

	return builder.String()
}

func (r TextRenderer) newRenderBuilder() *strings.Builder {
	var builder strings.Builder
	rowCount := r.Max.Y - r.Min.Y + 1
	columnCount := r.Max.X - r.Min.X + 1
	builder.Grow(rowCount*columnCount + rowCount - 1)
	return &builder
}

func (r TextRenderer) writeRow(builder *strings.Builder, cells Cells, y int) {
	for x := r.Min.X; x <= r.Max.X; x++ {
		state := cells[Point{X: x, Y: y}]
		cell := r.renderCell(state)
		builder.WriteRune(cell)
	}
}

func (r TextRenderer) renderCell(state State) rune {
	if state == Alive {
		return r.Alive
	}
	return r.Dead
}
