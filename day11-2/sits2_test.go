package sits2

import (
	"fmt"
	"strings"
	"testing"
)

func TestCanFindBoundaryStraightLeft(t *testing.T) {
	data := []string{"#......."}
	coord := Coordinate{X: 0, Y: 7, MaxIndexX: 0, MaxIndexY: 7}
	coord.FindCharacter(".", data, Straight, Left)
	if coord.identified != true || coord.X != 0 {
		t.Errorf("Could not find coordinate, received %+v", coord)
	}
}

func TestCanFindBoundaryStraightRight(t *testing.T) {
	data := []string{".......#"}
	coord := Coordinate{X: 0, Y: 0, MaxIndexX: 0, MaxIndexY: 7}
	coord.FindCharacter(".", data, Straight, Right)
	if coord.identified != true || coord.Y != 7 {
		t.Errorf("Could not find coordinate, received %+v", coord)
	}
}

func TestCanFindBoundaryStraightUp(t *testing.T) {
	data := []string{"#", ".", ".", "."}
	coord := Coordinate{X: 3, Y: 0, MaxIndexX: 3, MaxIndexY: 0}
	coord.FindCharacter(".", data, Up, Forward)
	if coord.identified != true || coord.X != 0 {
		t.Errorf("Could not find coordinate, received %+v", coord)
	}
}

func TestCanFindBoundaryStraightDown(t *testing.T) {
	data := []string{".", ".", ".", "#"}
	coord := Coordinate{X: 0, Y: 0, MaxIndexX: 3, MaxIndexY: 0}
	coord.FindCharacter(".", data, Down, Forward)
	if coord.identified != true || coord.X != 3 {
		t.Errorf("Could not find coordinate, received %+v", coord)
	}
}

func TestCanFindBoundaryUpperLeft(t *testing.T) {
	data := []string{"#..", "...", "..."}
	coord := Coordinate{X: 2, Y: 2, MaxIndexX: 2, MaxIndexY: 2}
	coord.FindCharacter(".", data, Up, Left)
	if coord.identified != true || coord.X != 0 || coord.Y != 0 {
		t.Errorf("Could not find coordinate, received %+v", coord)
	}
}

func TestCanFindBoundaryUpperRight(t *testing.T) {
	data := []string{"..#", "...", "..."}
	coord := Coordinate{X: 2, Y: 0, MaxIndexX: 2, MaxIndexY: 2}
	coord.FindCharacter(".", data, Up, Right)
	if coord.identified != true || coord.X != 0 || coord.Y != 2 {
		t.Errorf("Could not find coordinate, received %+v", coord)
	}
}

func TestCanFindBoundaryLowerLeft(t *testing.T) {
	data := []string{"...", "...", "#.."}
	coord := Coordinate{X: 0, Y: 2, MaxIndexX: 2, MaxIndexY: 2}
	coord.FindCharacter(".", data, Down, Left)
	if coord.identified != true || coord.X != 2 || coord.Y != 0 {
		t.Errorf("Could not find coordinate, received %+v", coord)
	}
}

func TestCanFindBoundaryLowerRight(t *testing.T) {
	data := []string{"...", "...", "..#"}
	coord := Coordinate{X: 0, Y: 0, MaxIndexX: 2, MaxIndexY: 2}
	coord.FindCharacter(".", data, Down, Right)
	if coord.identified != true || coord.X != 2 || coord.Y != 2 {
		t.Errorf("Could not find coordinate, received %+v", coord)
	}
}

func TestSearchBoxSeesToEndOfMap(t *testing.T) {
	data := Plane2([]string{
		".##.##.",
		"#.#.#.#",
		"##...##",
		"...L...",
		"##...##",
		"#.#.#.#",
		".##.##.",
	})
	home := Coordinate{X: 3, Y: 3, MaxIndexX: len(data) - 1, MaxIndexY: len(data[0]) - 1}
	sb := SearchBox{Home: home}
	sb.FindBoundaries(".", data)
	if sb.MinXMinY.X != 0 || sb.MinXMinY.Y != 0 || sb.MinXMidY.X != 0 || sb.MinXMidY.Y != 3 || sb.MinXMaxY.X != 0 || sb.MinXMaxY.Y != 6 {
		t.Errorf("First boundaries not correct")
	}
	if sb.MidXMinY.X != 3 || sb.MidXMinY.Y != 0 || sb.MidXMaxY.X != 3 || sb.MidXMaxY.Y != 6 {
		t.Errorf("Second boundaries not correct")
	}
	if sb.MaxXMinY.X != 6 || sb.MaxXMinY.Y != 0 || sb.MaxXMidY.X != 6 || sb.MaxXMidY.Y != 3 || sb.MaxXMaxY.X != 6 || sb.MaxXMaxY.Y != 6 {
		t.Errorf("Third boundaries are not correct")
	}
}

func TestSearchBoxSeesNothing(t *testing.T) {
	data := Plane2([]string{
		".##.##.",
		"#.#.#.#",
		"##...##",
		"...L...",
		"##...##",
		"#.#.#.#",
		".##.##.",
	})
	home := Coordinate{X: 3, Y: 3, MaxIndexX: len(data) - 1, MaxIndexY: len(data[0]) - 1}
	sb := SearchBox{Home: home}
	sb.FindBoundaries(".", data)
	countFound := sb.CountCharInBoundingBox("#", data)
	if countFound != 0 {
		t.Errorf("Could not find them")
	}
}

func TestSearchBoxSeesCorrectly(t *testing.T) {
	data := Plane2([]string{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	})
	home := Coordinate{X: 4, Y: 3, MaxIndexX: len(data) - 1, MaxIndexY: len(data[0]) - 1}
	sb := SearchBox{Home: home}
	sb.FindBoundaries(".", data)
	countFound := sb.CountCharInBoundingBox("#", data)
	if countFound != 8 {
		t.Errorf("Could not find them")
	}
}

func testPlanes(expected Plane2, actual Plane2) func(*testing.T) {
	return func(t *testing.T) {
		if strings.Join(actual, "") != strings.Join(expected, "") {
			t.Errorf("Expected \n%+v\nReceived \n%+v\n", strings.Join(expected, "\n"), strings.Join(actual, "\n"))
		}
	}
}

func TestAllPlanes(t *testing.T) {
	firstMap := Plane2([]string{
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
	secondMap := Plane2([]string{
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
	thirdMap := Plane2([]string{
		"#.LL.LL.L#",
		"#LLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLL#",
		"#.LLLLLL.L",
		"#.LLLLL.L#",
	})
	fourthMap := Plane2([]string{
		"#.L#.##.L#",
		"#L#####.LL",
		"L.#.#..#..",
		"##L#.##.##",
		"#.##.#L.##",
		"#.#####.#L",
		"..#.#.....",
		"LLL####LL#",
		"#.L#####.L",
		"#.L####.L#",
	})
	fifthMap := Plane2([]string{
		"#.L#.L#.L#",
		"#LLLLLL.LL",
		"L.L.L..#..",
		"##LL.LL.L#",
		"L.LL.LL.L#",
		"#.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLL#",
		"#.LLLLL#.L",
		"#.L#LL#.L#",
	})
	sixthMap := Plane2([]string{
		"#.L#.L#.L#",
		"#LLLLLL.LL",
		"L.L.L..#..",
		"##L#.#L.L#",
		"L.L#.#L.L#",
		"#.L####.LL",
		"..#.#.....",
		"LLL###LLL#",
		"#.LLLLL#.L",
		"#.L#LL#.L#",
	})
	seventhMap := Plane2([]string{
		"#.L#.L#.L#",
		"#LLLLLL.LL",
		"L.L.L..#..",
		"##L#.#L.L#",
		"L.L#.LL.L#",
		"#.LLLL#.LL",
		"..#.L.....",
		"LLL###LLL#",
		"#.LLLLL#.L",
		"#.L#LL#.L#",
	})
	fmt.Println(strings.Join(secondMap, "\n"))
	t.Run("1->2", testPlanes(secondMap, BuildPlane(firstMap)))
	t.Run("2->3", testPlanes(thirdMap, BuildPlane(secondMap)))
	t.Run("3->4", testPlanes(fourthMap, BuildPlane(thirdMap)))
	t.Run("4->5", testPlanes(fifthMap, BuildPlane(fourthMap)))
	t.Run("5->6", testPlanes(sixthMap, BuildPlane(fifthMap)))
	t.Run("6->7", testPlanes(seventhMap, BuildPlane(sixthMap)))
	t.Run("7->7", testPlanes(seventhMap, BuildPlane(seventhMap)))
}
