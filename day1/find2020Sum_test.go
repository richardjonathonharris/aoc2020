package find2020Sum

import (
	"testing"
)

func TestCanFindSumAtSpecificValue(t *testing.T) {
	testList := []int{1010, 1010}
	res, _ := FindFirst2ValuesEqualValue(testList, 2020)
	if res != [2]int{1010, 1010} {
		t.Errorf("FindFirst2ValuesEqualValue should return [1010, 1010], returned %d", res)
	}
}

func TestRaisesErrorIfSumNotFound(t *testing.T) {
	testList := []int{1, 2}
	_, error := FindFirst2ValuesEqualValue(testList, 2020)
	if error == nil {
		t.Errorf("FindFirst2ValuesEqualValue should raise error if no sum can be found")
	}
}

func TestFindsFirstPairOfSumsToMeetCriteria(t *testing.T) {
	testList := []int{1010, 1010, 1, 2019}
	res, _ := FindFirst2ValuesEqualValue(testList, 2020)
	if res == [2]int{1, 2019} {
		t.Errorf("FindFirst2ValuesEqualValue should return first pair [1010, 1010], returned %d", res)
	}
}

func TestFindsProductAt3(t *testing.T) {
	testList := []int{1010, 1010, 1, 2019}
	res, _ := FindMultipleEqualValue(testList, 2021, 3)
	if res != 1010*1010 {
		t.Errorf("Should find the product for the first 3 terms that sum to 2021")
	}
}
