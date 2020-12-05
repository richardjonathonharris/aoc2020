package seats

import (
	"fmt"
	"testing"
)

func TestCanReturnRowIndicesBack(t *testing.T) {
	resp := SplitSeats("B", [2]int{0, 127})
	if resp != [2]int{64, 127} {
		t.Errorf("Cannot split rows in back to {64, 127}")
	}
}

func TestCanReturnRowIndicesFront(t *testing.T) {
	resp := SplitSeats("F", [2]int{0, 127})
	if resp != [2]int{0, 63} {
		t.Errorf("Cannot split rows in front to {0, 63}")
	}
}

func TestCanReturnColumnIndicesRight(t *testing.T) {
	resp := SplitSeats("R", [2]int{0, 7})
	if resp != [2]int{4, 7} {
		t.Errorf("Cannot split seats to right to {4, 7}")
	}
}

func TestCanReturnColumnIndicesLeft(t *testing.T) {
	resp := SplitSeats("L", [2]int{0, 7})
	if resp != [2]int{0, 3} {
		t.Errorf("Cannot split seats to left to {0, 3}")
	}
}

func testValidators(expectedRow int, expectedSeat int, expectedSeatId int, actualRow int, actualSeat int, actualSeatId int) func(*testing.T) {
	return func(t *testing.T) {
		if actualRow != expectedRow {
			t.Errorf(fmt.Sprintf("Expected %d but got %d for row", expectedRow, actualRow))
		}
		if actualSeat != expectedSeat {
			t.Errorf(fmt.Sprintf("Expected %d but got %d for seat", expectedSeat, actualSeat))
		}
		if actualSeatId != expectedSeatId {
			t.Errorf(fmt.Sprintf("Expected %d but got %d for seat Id", expectedSeatId, actualSeatId))
		}
	}
}

func TestAllExamples(t *testing.T) {
	cp := FindSeats("FBFBBFFRLR", [2]int{0, 127}, [2]int{0, 7})
	row, column, id, error := cp.GetSeatInformation()
	if error != nil {
		t.Errorf("Did not work")
	}
	t.Run("FBFBBFFRLR", testValidators(44, 5, 357, row, column, id))

	cp1 := FindSeats("BFFFBBFRRR", [2]int{0, 127}, [2]int{0, 7})
	row1, column1, id1, error1 := cp1.GetSeatInformation()
	if error1 != nil {
		t.Errorf("Did not work")
	}
	t.Run("BFFFBBFRRR", testValidators(70, 7, 567, row1, column1, id1))

	cp2 := FindSeats("FFFBBBFRRR", [2]int{0, 127}, [2]int{0, 7})
	row2, column2, id2, error2 := cp2.GetSeatInformation()
	if error2 != nil {
		t.Errorf("Did not work")
	}
	t.Run("FFFBBBFRRR", testValidators(14, 7, 119, row2, column2, id2))

	cp3 := FindSeats("BBFFBBFRLL", [2]int{0, 127}, [2]int{0, 7})
	row3, column3, id3, error3 := cp3.GetSeatInformation()
	if error3 != nil {
		t.Errorf("Did not work")
	}
	t.Run("BBFFBBFRLL", testValidators(102, 4, 820, row3, column3, id3))
}
