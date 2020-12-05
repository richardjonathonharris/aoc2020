package seats

import "errors"

type CurrentPosition struct {
	row    [2]int
	column [2]int
}

func (cp *CurrentPosition) GetSeatInformation() (int, int, int, error) {
	if (cp.row[0] != cp.row[1]) || (cp.column[0] != cp.column[1]) {
		return 0, 0, 0, errors.New("Cannot determine seat")
	} else {
		return cp.row[0], cp.column[0], (cp.row[0] * 8) + cp.column[0], nil
	}
}

func SplitSeats(direction string, currentBoundary [2]int) [2]int {
	numRows := currentBoundary[1] - currentBoundary[0]
	splitRow := currentBoundary[0] + (numRows / 2)
	if direction == "F" || direction == "L" {
		return [2]int{currentBoundary[0], splitRow}
	} else if direction == "B" || direction == "R" {
		return [2]int{splitRow + 1, currentBoundary[1]}
	} else {
		return currentBoundary
	}
}

func FindSeats(direction string, startingRowBox [2]int, startingColumnBox [2]int) CurrentPosition {
	currentPosition := CurrentPosition{row: startingRowBox, column: startingColumnBox}
	rows := direction[0:7]
	seats := direction[7:10]
	for _, rowDirection := range rows {
		currentPosition.row = SplitSeats(string(rowDirection), currentPosition.row)
	}
	for _, seatDirection := range seats {
		currentPosition.column = SplitSeats(string(seatDirection), currentPosition.column)
	}
	return currentPosition
}
