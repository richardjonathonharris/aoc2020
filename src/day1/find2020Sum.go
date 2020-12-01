package day1

import "errors"

func FindFirst2ValuesEqualValue(listOfEntries []int, sumTotal int) ([2]int, error) {
	for idx, value := range listOfEntries {
		for innerIdx, pairValue := range listOfEntries {
			if idx == innerIdx {
				continue
			}
			summed := value + pairValue
			if summed == sumTotal {
				return [2]int{value, pairValue}, nil
			} else {
				continue
			}
		}
	}
	return [2]int{0,0}, errors.New("Could not find pair in list that summed to total")
}