package game

import (
	"testing"
)

func testNumberGame(expected int, actual int) func(*testing.T) {
	return func(t *testing.T) {
		if expected != actual {
			t.Errorf("Should have value as %d, received %d", expected, actual)
		}
	}
}

func TestGames(t *testing.T) {
	data0 := []int{0, 3, 6}
	data1 := []int{1, 3, 2}
	data2 := []int{2, 1, 3}
	data3 := []int{1, 2, 3}
	data4 := []int{2, 3, 1}
	data5 := []int{3, 2, 1}
	data6 := []int{3, 1, 2}
	t.Run("0", testNumberGame(0, NumberGame(data0, 10)))
	t.Run("1", testNumberGame(1, NumberGame(data1, 2020)))
	t.Run("2", testNumberGame(10, NumberGame(data2, 2020)))
	t.Run("3", testNumberGame(27, NumberGame(data3, 2020)))
	t.Run("4", testNumberGame(78, NumberGame(data4, 2020)))
	t.Run("5", testNumberGame(438, NumberGame(data5, 2020)))
	t.Run("6", testNumberGame(1836, NumberGame(data6, 2020)))
}
