package xmas

import (
	"fmt"
	"testing"
)

func TestIdentifyInvalidValueNotPairOfPastValues(t *testing.T) {
	preamble := []int{1, 2, 3, 4, 5}
	checkValue := 100
	resp := IdentifyValidValue(preamble, checkValue)
	if resp != false {
		t.Errorf("Value is not sum of two numbers in preamble")
	}
}

func TestIdentifyInvalidValuePairOfSamePastValue(t *testing.T) {
	preamble := []int{1, 2, 3, 4, 4, 5}
	checkValue := 8
	resp := IdentifyValidValue(preamble, checkValue)
	if resp != false {
		t.Errorf("Value is sum of two of the same value")
	}
}

func TestIdentifyValidValue(t *testing.T) {
	preamble := []int{1, 2, 3, 4, 5}
	checkValue := 9
	resp := IdentifyValidValue(preamble, checkValue)
	if resp != true {
		t.Errorf("Value is sum of two numbers in preamble AND is not sum of two of the same value")
	}
}

func TestFindFirstInvalidValue(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 12}
	resp, err := FindFirstInvalidIndex(data, 3, 3)
	if err != nil && resp != 7 {
		t.Errorf("Did not find correct invalid value")
	}
}

func TestFindSmallLargeSumValues(t *testing.T) {
	data := []int{1, 2, 3, 3, 6, 12}
	minVal, maxVal := FindSmallLargeSumValues(data, 5)
	fmt.Println(minVal, maxVal)
	if minVal != 3 && maxVal != 6 {
		t.Errorf("Did not find earliest and latest value")
	}
}
