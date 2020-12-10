package joltage

import (
	"errors"
	"fmt"
)

type JoltDiff int

const (
	One   JoltDiff = 1
	Two            = 2
	Three          = 3
)

type JoltRecord map[string]int

func GetNextJoltDiff(currJolt int, nextJolt int) (JoltDiff, error) {
	if nextJolt - currJolt < 1 || nextJolt - currJolt > 3 {
		return One, errors.New("Jolt Difference is too big")
	}
	return JoltDiff(nextJolt - currJolt), nil
}

func GetPossibleNextJoltDiffs(currJolt int, data[]int) []int {
	goodJolts := []int{}
	for _, nextJolt := range data {
		if nextJolt - currJolt < 1 || nextJolt - currJolt > 3 {
			continue
		} else {
			goodJolts = append(goodJolts, nextJolt)
		}
	}
	return goodJolts
}

func DeterminePotentialRoutes(records *[][]int, currentRecord []int, data[]int, maxValue int) {
	if len(*records) % 1000  == 0 {
		fmt.Println("Currently at ", len(*records))
	}
	for i := 1; i < 4; i++ {
		valueInData := false
		newValue := currentRecord[len(currentRecord) - 1] + i
		for _, d := range data {
			if d == newValue {
				valueInData = true
			}
		}
		if newValue > maxValue || !valueInData {
			continue
		}
		newRecord := append(currentRecord, newValue) 
		if newValue == maxValue {
			*records = append(*records, newRecord)
		} else {
			DeterminePotentialRoutes(records, newRecord, data, maxValue)
		}
	}
}

func GetMapOfJoltDiffs(data []int) (JoltRecord, error) {
	jolts := JoltRecord{"1": 0, "2": 0, "3": 0}
	for idx, jolt := range data {
		if idx == len(data) -1 {
			break
		}
		joltDiff, err := GetNextJoltDiff(jolt, data[idx + 1])
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