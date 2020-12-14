package bus

import (
	"testing"
)

func TestGetRemainder(t *testing.T) {
	resp := FindRemainder(944, 939)
	if resp != 5 {
		t.Errorf("We're just testing that modulo works as we expect, received %d expected 5", resp)
	}
}

func TestFindClosestTimestamp(t *testing.T) {
	resp := FindClosestTimestamp(939, 59)
	if resp != 944 {
		t.Errorf("Should find that the closest departure to 939 is 944, got %d", resp)
	}
}

func TestFindBestOption(t *testing.T) {
	data := []string{"7", "13", "x", "x", "59", "x", "31", "19"}
	resp := SolveDay1(939, data)
	if resp != 295 {
		t.Errorf("Day 1 example not solved correctly, got %d", resp)
	}
}

func TestFindOffsets(t *testing.T) {
	data := []string{"7", "13", "x", "x", "59", "x", "31", "19"}
	resp := GetRequiredOffsets(data)
	if resp[0] != 7 || resp[1] != 13 || resp[4] != 59 || resp[6] != 31 || resp[7] != 19 {
		t.Errorf("Should have gotten properly constructed set of offsets %+v", resp)
	}
}

func testOffset(expected int, actual int) func(*testing.T) {
	return func(t *testing.T) {
		if expected != actual {
			t.Errorf("Should have first timestamp as %d, received %d", expected, actual)
		}
	}
}

func TestOffsets(t *testing.T) {
	data1 := []string{"17", "x", "13", "19"}
	data2 := []string{"67", "7", "59", "61"}
	data3 := []string{"67", "x", "7", "59", "61"}
	data4 := []string{"67", "7", "x", "59", "61"}
	data5 := []string{"1789", "37", "47", "1889"}
	t.Run("1", testOffset(3417, FindEarliestTimestampsInline(data1)))
	t.Run("2", testOffset(754018, FindEarliestTimestampsInline(data2)))
	t.Run("3", testOffset(779210, FindEarliestTimestampsInline(data3)))
	t.Run("4", testOffset(1261476, FindEarliestTimestampsInline(data4)))
	t.Run("5", testOffset(1202161486, FindEarliestTimestampsInline(data5)))
}
