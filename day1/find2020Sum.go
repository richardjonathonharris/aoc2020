package find2020Sum

import (
	"errors"
	"gonum.org/v1/gonum/stat/combin"
)

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

func FindMultipleEqualValue(listOfEntries []int, sumTotal int, numValues int) (int, error) {
	// Find all unique combinations of numValues size
	combinations := combin.Combinations(len(listOfEntries), numValues)
	var summedValues []int
	var productValues []int
	for _, combo := range combinations {
		summedValue := 0
		product := 1
		for _, idx := range combo {
			summedValue += listOfEntries[idx]
			product *= listOfEntries[idx]
		}
		summedValues = append(summedValues, summedValue)
		productValues = append(productValues, product)
	}
	for i, val := range summedValues {
		if val == sumTotal {
			return productValues[i], nil
		}
	}
	return 0, errors.New("Could not find combination whose sum totaled")
}