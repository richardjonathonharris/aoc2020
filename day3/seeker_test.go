package seeker

import (
	"testing"
)

func TestCanFindNextIndexOnInstruction(t *testing.T) {
	direct := Direction{Right: 2, Down: 1}
	startingPosition := Index{Row: 0, Column: 0}
	resp, _ := GetNextIndex(direct, startingPosition, 10, 10)
	newPosition := Index{Row: 2, Column: 1}
	if resp != newPosition {
		t.Errorf("GetNextIndex should return row 2, column 1")
	}
}

func TestCanLoopOverRightSideOfHill(t *testing.T) {
	direct := Direction{Right: 2, Down: 1}
	startingPosition := Index{Row: 0, Column: 0}
	resp, _ := GetNextIndex(direct, startingPosition, 1, 10)
	newPosition := Index{Row: 0, Column: 1}
	if resp != newPosition {
		t.Errorf("GetNextIndex should return row 0, column 1")
	}
}

func TestCanLoopOverLeftSideOfHill(t *testing.T) {
	direct := Direction{Left: 2, Down: 1}
	startingPosition := Index{Row: 0, Column: 0}
	resp, _ := GetNextIndex(direct, startingPosition, 2, 10)
	newPosition := Index{Row: 0, Column: 1}
	if resp != newPosition {
		t.Errorf("GetNextIndex should return row 0, column 1")
	}
}

func TestRaisesWhenAtBottomOfHill(t *testing.T) {
	direct := Direction{Down: 10}
	startingPosition := Index{Row: 0, Column: 0}
	_, err := GetNextIndex(direct, startingPosition, 1, 1)
	if err == nil {
		t.Errorf("Should raise error when at bottom of hill")
	}
}

func TestRaisesWhenAtTopOfHill(t *testing.T) {
	direct := Direction{Up: 10}
	startingPosition := Index{Row: 0, Column: 0}
	_, err := GetNextIndex(direct, startingPosition, 10, 10)
	if err == nil {
		t.Errorf("Should raise error when at top of hill")
	}
}
