package sits

import (
	"strings"
	"testing"
)

func TestFindSurroundingSpots(t *testing.T) {
	sb := SurroundingBox{X: 1, Y: 1}
	sb.FindSurroundingBox(2, 2)
	expectedSb := SurroundingBox{X: 1, Y: 1, XMin: 0, XMax: 2, YMin: 0, YMax: 2}
	if sb != expectedSb {
		t.Errorf("Expected %+v, got %+v", expectedSb, sb)
	}
}

func TestHandlesBoxOnCorner(t *testing.T) {
	sb := SurroundingBox{X: 0, Y: 0}
	sb.FindSurroundingBox(2, 2)
	expectedSb := SurroundingBox{X: 0, Y: 0, XMin: 0, XMax: 1, YMin: 0, YMax: 1}
	if sb != expectedSb {
		t.Errorf("Expected %+v, got %+v", expectedSb, sb)
	}
}

func TestHandlesBoxOnOppositeCorner(t *testing.T) {
	sb := SurroundingBox{X: 2, Y: 2}
	sb.FindSurroundingBox(2, 2)
	expectedSb := SurroundingBox{X: 2, Y: 2, XMin: 1, XMax: 2, YMin: 1, YMax: 2}
	if sb != expectedSb {
		t.Errorf("Expected %+v, got %+v", expectedSb, sb)
	}
}

func TestCountsCharInCenter(t *testing.T) {
	sb := SurroundingBox{X: 1, Y: 1}
	sb.FindSurroundingBox(2, 2)
	data := []string{"###", ".L.", "..."}
	data = Plane(data)
	count := CountCharInBoundingBox("#", data, sb)
	if count != 3 {
		t.Errorf("Should count 3 # nearby, counted %d", count)
	}
}

func TestCountsCharOnCorner(t *testing.T) {
	sb := SurroundingBox{X: 0, Y: 0}
	sb.FindSurroundingBox(2, 2)
	data := []string{"###", ".L.", "..."}
	data = Plane(data)
	count := CountCharInBoundingBox("#", data, sb)
	if count != 1 {
		t.Errorf("Should count 2 # nearby, counted %d", count)
	}
}

func testPlanes(expected Plane, actual Plane) func(*testing.T) {
	return func(t *testing.T) {
		if strings.Join(actual, "") != strings.Join(expected, "") {
			t.Errorf("Expected \n%+v\nReceived \n%+v\n", strings.Join(expected, "\n"), strings.Join(actual, "\n"))
		}
	}
}

func TestAllPlanes(t *testing.T) {
	firstMap := Plane([]string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	})
	secondMap := Plane([]string{
		"#.##.##.##",
		"#######.##",
		"#.#.#..#..",
		"####.##.##",
		"#.##.##.##",
		"#.#####.##",
		"..#.#.....",
		"##########",
		"#.######.#",
		"#.#####.##",
	})
	thirdMap := Plane([]string{
		"#.LL.L#.##",
		"#LLLLLL.L#",
		"L.L.L..L..",
		"#LLL.LL.L#",
		"#.LL.LL.LL",
		"#.LLLL#.##",
		"..L.L.....",
		"#LLLLLLLL#",
		"#.LLLLLL.L",
		"#.#LLLL.##",
	})
	fourthMap := Plane([]string{
		"#.##.L#.##",
		"#L###LL.L#",
		"L.#.#..#..",
		"#L##.##.L#",
		"#.##.LL.LL",
		"#.###L#.##",
		"..#.#.....",
		"#L######L#",
		"#.LL###L.L",
		"#.#L###.##",
	})
	fifthMap := Plane([]string{
		"#.#L.L#.##",
		"#LLL#LL.L#",
		"L.L.L..#..",
		"#LLL.##.L#",
		"#.LL.LL.LL",
		"#.LL#L#.##",
		"..L.L.....",
		"#L#LLLL#L#",
		"#.LLLLLL.L",
		"#.#L#L#.##",
	})
	sixthMap := Plane([]string{
		"#.#L.L#.##",
		"#LLL#LL.L#",
		"L.#.L..#..",
		"#L##.##.L#",
		"#.#L.LL.LL",
		"#.#L#L#.##",
		"..L.L.....",
		"#L#L##L#L#",
		"#.LLLLLL.L",
		"#.#L#L#.##",
	})
	t.Run("1->2", testPlanes(secondMap, BuildPlane(firstMap)))
	t.Run("2->3", testPlanes(thirdMap, BuildPlane(secondMap)))
	t.Run("3->4", testPlanes(fourthMap, BuildPlane(thirdMap)))
	t.Run("4->5", testPlanes(fifthMap, BuildPlane(fourthMap)))
	t.Run("5->6", testPlanes(sixthMap, BuildPlane(fifthMap)))
	t.Run("6->6", testPlanes(sixthMap, BuildPlane(sixthMap)))
}
