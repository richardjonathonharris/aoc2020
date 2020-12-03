package seeker

import (
	"errors"
	"fmt"
	"math"
)

type Direction struct {
	Up    int
	Down  int
	Left  int
	Right int
}

type Index struct {
	Row    int
	Column int
}

func overflowRow(currentIndex int, desiredDirection int, maxRow int) int {
	newIndex := currentIndex + desiredDirection
	if newIndex < 0 {
		// loops around to right side
		return maxRow + newIndex
	} else if newIndex > maxRow {
		// loops past right side back over to left (we lose one flipping over)
		return newIndex - maxRow - 1
	} else {
		return newIndex
	}
}

func GetNextIndex(direction Direction, currentIndex Index, maxRow int, maxCol int) (Index, error) {
	desiredRowDirection := direction.Right - direction.Left
	newRowIndex := overflowRow(currentIndex.Row, desiredRowDirection, maxRow)
	newColumnIndex := currentIndex.Column + direction.Down - direction.Up
	if newColumnIndex < 0 || newColumnIndex > maxCol {
		return Index{}, errors.New("Directions would take user above or below end of path")
	} else {
		return Index{Row: newRowIndex, Column: newColumnIndex}, nil
	}
}

func Day3HelperFunction(direction Direction, data []string, treeValue string) int {
	numRowsToTraverse := len(data)
	runLoop := math.Ceil(float64(numRowsToTraverse) / float64(direction.Down))
	indices := []Index{{}}
	trees := 0
	for i := 0; i < int(runLoop); i++ {
		nextIndex, err := GetNextIndex(direction, indices[len(indices)-1], len(data[0])-1, numRowsToTraverse-1)
		if err != nil {
			fmt.Println("We reached the end of the slope")
			break
		}
		if string(data[nextIndex.Column][nextIndex.Row]) == string(treeValue) {
			trees += 1
		}
		indices = append(indices, nextIndex)
	}
	return trees
}
