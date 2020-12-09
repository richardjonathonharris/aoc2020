package xmas

import (
	"errors"
)

func IdentifyValidValue(preamble []int, checkValue int) bool {
	foundPair := false
	foundPairWithDouble := false
	for idx1, val1 := range preamble {
		for idx2, val2 := range preamble {
			if idx1 == idx2 {
				continue
			}
			if (val1 != val2) && (val1+val2 == checkValue) {
				foundPair = true
			}
			if (val1 == val2) && (val1+val2 == checkValue) {
				foundPairWithDouble = true
			}
		}
	}
	if foundPairWithDouble {
		return false
	} else {
		return foundPair
	}
}

func FindFirstInvalidIndex(data []int, windowSize int, startingIndex int) (int, error) {
	for i := startingIndex; i < len(data); i++ {
		if !IdentifyValidValue(data[i-windowSize:i], data[i]) {
			return i, nil
		}
	}
	return 0, errors.New("Could not find an invalid value")
}

func FindSmallLargeSumValues(data []int, badIndex int) (int, int) {
	for j := 0; j < len(data[0:badIndex]); j++ {
		modifiedBadIndex := badIndex - j
		for i := 0; i < len(data[0:modifiedBadIndex]); i++ {
			currSum := 0
			for _, v := range data[modifiedBadIndex-i : modifiedBadIndex] {
				currSum += v
			}
			if currSum == data[badIndex] {
				smallestVal := data[modifiedBadIndex-i]
				biggestVal := data[modifiedBadIndex-1]
				for _, val := range data[modifiedBadIndex-i : modifiedBadIndex] {
					if val > biggestVal {
						biggestVal = val
					}
					if val < smallestVal {
						smallestVal = val
					}
				}
				return smallestVal, biggestVal
			} else if currSum > data[badIndex] {
				continue
			}
		}
	}
	return 0, 0
}
