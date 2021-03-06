package joltage

import (
	"fmt"
	"testing"
)

func TestCanFindJoltDifference(t *testing.T) {
	resp, err := GetNextJoltDiff(2, 5)
	if err != nil || resp != Three {
		t.Errorf("Should return three for next jolt difference.")
	}
}

func TestRaisesErrorOnInvalidComparison(t *testing.T) {
	_, err := GetNextJoltDiff(0, 0)
	if err == nil {
		t.Errorf("Should return error on invalid comparison")
	}
}

func TestCanCreateMapOfJoltDiffs(t *testing.T) {
	data := []int{1, 2, 4, 5}
	resp, err := GetMapOfJoltDiffs(data)
	if err != nil {
		t.Errorf("Got an unexpected error")
	}
	if resp["1"] != 2 || resp["2"] != 1 || resp["3"] != 0 {
		t.Errorf("Did not assign values as expected")
	}
}

func TestGetArrayOfNextJoltDiffs(t *testing.T) {
	data := []int{0, 2, 3, 4, 5}
	resp := GetPossibleNextJoltDiffs(1, data)
	if len(resp) != len([]int{2, 3, 4}) {
		t.Errorf("Did not get all possible values")
	}
}

func TestDynamicProgrammingOption(t *testing.T) {
	data := []int{0, 2, 3, 4, 5}
	resp := DynamicProgrammingOption(0, data, make(map[int]int))
	fmt.Println(resp)
	if resp != 4 {
		t.Errorf("Didn't find the right value")
	}
}
