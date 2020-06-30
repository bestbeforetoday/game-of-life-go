package life

import (
	"fmt"
	"testing"
)

func TestLiveCell(t *testing.T) {
	neighboursStateMap := map[int]State{
		0: Dead,
		1: Dead,
		2: Alive,
		3: Alive,
		4: Dead,
		5: Dead,
		6: Dead,
		7: Dead,
		8: Dead,
	}

	for neighbourCount, nextState := range neighboursStateMap {
		cell := Point{1, 1}
		neighbours := neighbours(cell)[0:neighbourCount]
		livesText := "Dies"
		if nextState == Alive {
			livesText = "Survives"
		}

		t.Run(fmt.Sprintf("%s with %d neighbours", livesText, len(neighbours)), func(t *testing.T) {
			game := NewGame(append(neighbours, cell))
			result := game.Next()
			if result.cells[cell] != nextState {
				t.Error()
			}
		})
	}
}

func TestDeadCell(t *testing.T) {
	neighboursStateMap := map[int]State{
		0: Dead,
		1: Dead,
		2: Dead,
		3: Alive,
		4: Dead,
		5: Dead,
		6: Dead,
		7: Dead,
		8: Dead,
	}

	for neighbourCount, nextState := range neighboursStateMap {
		cell := Point{1, 1}
		neighbours := neighbours(cell)[0:neighbourCount]
		livesText := "Stays dead"
		if nextState == Alive {
			livesText = "Becomes live"
		}

		t.Run(fmt.Sprintf("%s with %d neighbours", livesText, len(neighbours)), func(t *testing.T) {
			game := NewGame(neighbours)
			result := game.Next()
			if result.cells[cell] != nextState {
				t.Error()
			}
		})
	}
}
