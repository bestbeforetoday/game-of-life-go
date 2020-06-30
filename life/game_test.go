package life

import (
	"fmt"
	"testing"
)

func TestLiveCell(t *testing.T) {
	neighborsStateMap := map[int]State{
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

	for neighborCount, nextState := range neighborsStateMap {
		cell := Point{1, 1}
		neighbors := unboundedNeighbors(cell)[0:neighborCount]
		livesText := "Dies"
		if nextState == Alive {
			livesText = "Survives"
		}

		t.Run(fmt.Sprintf("%s with %d neighbors", livesText, len(neighbors)), func(t *testing.T) {
			game := NewUnboundedGame(append(neighbors, cell))
			result := game.Next()
			if result.cells[cell] != nextState {
				t.Error()
			}
		})
	}
}

func TestDeadCell(t *testing.T) {
	neighborsStateMap := map[int]State{
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

	for neighborCount, nextState := range neighborsStateMap {
		cell := Point{1, 1}
		neighbors := unboundedNeighbors(cell)[0:neighborCount]
		livesText := "Stays dead"
		if nextState == Alive {
			livesText = "Becomes live"
		}

		t.Run(fmt.Sprintf("%s with %d neighbors", livesText, len(neighbors)), func(t *testing.T) {
			game := NewUnboundedGame(neighbors)
			result := game.Next()
			if result.cells[cell] != nextState {
				t.Error()
			}
		})
	}
}

func TestBoundaries(t *testing.T) {
	t.Run("Propellor runs without boundary", func(t *testing.T) {
		cells := []Point{
			{-1, 0},
			{0, 0},
			{1, 0},
		}
		game := NewUnboundedGame(cells)

		result := game.Next().Next().Cells()

		if !hasExactlyLocations(result, cells) {
			t.Errorf("Expected %v, got %v", cells, result)
		}
	})

	t.Run("Propellor centered on (1,0) does not run with (0,0) boundary", func(t *testing.T) {
		cells := []Point{
			{0, 0},
			{1, 0},
			{2, 0},
		}
		game := NewBoundedGame(cells, Point{0, 0}, Point{2, 1})

		firstIteration := game.Next()

		expected1 := []Point{
			{1, 0},
			{1, 1},
		}
		if !hasExactlyLocations(firstIteration.Cells(), expected1) {
			t.Errorf("First iteration expected %v, got %v", expected1, firstIteration.Cells())
		}

		secondIteration := firstIteration.Next()

		expected2 := []Point{}
		if !hasExactlyLocations(secondIteration.Cells(), expected2) {
			t.Errorf("Second iteration expected %v, got %v", expected2, secondIteration.Cells())
		}

	})
}

func hasExactlyLocations(cells Cells, locations []Point) bool {
	if len(cells) != len(locations) {
		return false
	}

	for _, location := range locations {
		if cells[location] != Alive {
			return false
		}
	}

	return true
}
