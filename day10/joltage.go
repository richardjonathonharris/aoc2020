package joltage

import (
	"errors"
)

type JoltDiff int

const (
	One   JoltDiff = 1
	Two            = 2
	Three          = 3
)

type JoltRecord map[string]int

func GetNextJoltDiff(currJolt int, nextJolt int) (JoltDiff, error) {
	if nextJolt-currJolt < 1 || nextJolt-currJolt > 3 {
		return One, errors.New("Jolt Difference is too big")
	}
	return JoltDiff(nextJolt - currJolt), nil
}

func GetPossibleNextJoltDiffs(currJolt int, data []int) []int {
	goodJolts := []int{}
	for _, nextJolt := range data {
		if nextJolt-currJolt < 1 || nextJolt-currJolt > 3 {
			continue
		} else {
			goodJolts = append(goodJolts, nextJolt)
		}
	}
	return goodJolts
}

func GetMapOfJoltDiffs(data []int) (JoltRecord, error) {
	jolts := JoltRecord{"1": 0, "2": 0, "3": 0}
	for idx, jolt := range data {
		if idx == len(data)-1 {
			break
		}
		joltDiff, err := GetNextJoltDiff(jolt, data[idx+1])
		if err != nil {
			return jolts, errors.New("Couldn't create map because invalid jolt diff")
		}
		switch joltDiff {
		case One:
			jolts["1"] += 1
		case Two:
			jolts["2"] += 1
		case Three:
			jolts["3"] += 1
		}
	}
	return jolts, nil
}

func DynamicProgrammingOption(startIndex int, numbers []int, visited map[int]int) int {
	// shamelessly stolen: https://github.com/j4rv/advent-of-code-2020/blob/main/day-10/main.go

	// Because this recurses, if we're at the end of the array, there's only one more place left to go
	if startIndex >= len(numbers)-3 {
		return 1
	}

	currentNumber := numbers[startIndex]
	// if we've visited this index
	if result, ok := visited[currentNumber]; ok {
		return result
	}

	var count int
	// iterate for the next three indices
	for i := startIndex + 1; i < startIndex+4; i++ {
		num := numbers[i]
		if areCompatible(currentNumber, num) {
			count += DynamicProgrammingOption(i, numbers, visited)
		}
	}

	// Find number of counts of the current number
	visited[currentNumber] = count
	return count
}

func areCompatible(low int, high int) bool {
	return low+1 == high || low+2 == high || low+3 == high
}
