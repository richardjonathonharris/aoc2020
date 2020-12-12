package sits2

import (
	"fmt"
)

type Plane2 []string

type Vertical string

const (
	Up       Vertical = "up"
	Straight          = "straight"
	Down              = "down"
)

type Horizontal string

const (
	Left    Horizontal = "left"
	Forward            = "forward"
	Right              = "right"
)

type Coordinate struct {
	// Some coordinate on a Plane, Y traverses within a row, so
	// . . .
	// . . X
	// the X would be coordinate {Y: 3, X: 2}
	X          int
	Y          int
	identified bool
	MaxIndexX  int
	MaxIndexY  int
}

func (c *Coordinate) Print() {
	fmt.Printf("X: %d, Y: %d\n", c.X, c.Y)
}

type SearchBox struct {
	// We need to hold what coordinate to look at in 8 directions.
	// Each of these is a coordinate, so for a plane like:
	// 1 2 3
	// 4 * 5
	// 6 7 8
	// 1 == MinXMinY
	// 3 == MinXMaxY
	// 6 == MaxXMinY
	// 8 == MaxXMaxY
	// * == Home
	Home     Coordinate
	MinXMinY Coordinate
	MinXMidY Coordinate
	MinXMaxY Coordinate
	MidXMinY Coordinate
	MidXMaxY Coordinate
	MaxXMinY Coordinate
	MaxXMidY Coordinate
	MaxXMaxY Coordinate
}

func (sb *SearchBox) FindBoundaries(charac string, plane Plane2) {
	// initialize starting positions
	var effectiveMinX, effectiveMinY, effectiveMaxX, effectiveMaxY int
	if sb.Home.X <= 0 {
		effectiveMinX = 0
	} else {
		effectiveMinX = sb.Home.X - 1
	}
	if sb.Home.X >= sb.Home.MaxIndexX {
		effectiveMaxX = sb.Home.X
	} else {
		effectiveMaxX = sb.Home.X + 1
	}
	if sb.Home.Y <= 0 {
		effectiveMinY = 0
	} else {
		effectiveMinY = sb.Home.Y - 1
	}
	if sb.Home.Y >= sb.Home.MaxIndexY {
		effectiveMaxY = sb.Home.Y
	} else {
		effectiveMaxY = sb.Home.Y + 1
	}
	sb.MinXMinY = Coordinate{X: effectiveMinX, Y: effectiveMinY, MaxIndexX: sb.Home.MaxIndexX, MaxIndexY: sb.Home.MaxIndexY}
	sb.MinXMidY = Coordinate{X: effectiveMinX, Y: sb.Home.Y, MaxIndexX: sb.Home.MaxIndexX, MaxIndexY: sb.Home.MaxIndexY}
	sb.MinXMaxY = Coordinate{X: effectiveMinX, Y: effectiveMaxY, MaxIndexX: sb.Home.MaxIndexX, MaxIndexY: sb.Home.MaxIndexY}
	sb.MidXMinY = Coordinate{X: sb.Home.X, Y: effectiveMinY, MaxIndexX: sb.Home.MaxIndexX, MaxIndexY: sb.Home.MaxIndexY}
	sb.MidXMaxY = Coordinate{X: sb.Home.X, Y: effectiveMaxY, MaxIndexX: sb.Home.MaxIndexX, MaxIndexY: sb.Home.MaxIndexY}
	sb.MaxXMinY = Coordinate{X: effectiveMaxX, Y: effectiveMinY, MaxIndexX: sb.Home.MaxIndexX, MaxIndexY: sb.Home.MaxIndexY}
	sb.MaxXMidY = Coordinate{X: effectiveMaxX, Y: sb.Home.Y, MaxIndexX: sb.Home.MaxIndexX, MaxIndexY: sb.Home.MaxIndexY}
	sb.MaxXMaxY = Coordinate{X: effectiveMaxX, Y: effectiveMaxY, MaxIndexX: sb.Home.MaxIndexX, MaxIndexY: sb.Home.MaxIndexY}

	// Then search out to find coordinates
	sb.MinXMinY.FindCharacter(charac, plane, Up, Left)
	sb.MinXMidY.FindCharacter(charac, plane, Up, Forward)
	sb.MinXMaxY.FindCharacter(charac, plane, Up, Right)
	sb.MidXMinY.FindCharacter(charac, plane, Straight, Left)
	sb.MidXMaxY.FindCharacter(charac, plane, Straight, Right)
	sb.MaxXMinY.FindCharacter(charac, plane, Down, Left)
	sb.MaxXMidY.FindCharacter(charac, plane, Down, Forward)
	sb.MaxXMaxY.FindCharacter(charac, plane, Down, Right)
}

func (sb *SearchBox) ListOfCoordinates() []Coordinate {
	return []Coordinate{
		sb.MinXMinY, sb.MinXMidY, sb.MinXMaxY,
		sb.MidXMinY, sb.MidXMaxY,
		sb.MaxXMinY, sb.MaxXMidY, sb.MaxXMaxY,
	}
}

func (sb *SearchBox) CountCharInBoundingBox(charac string, plane Plane2) int {
	count := 0
	coordinatesSeen := make(map[string]int)
	for _, coord := range sb.ListOfCoordinates() {
		coordString := fmt.Sprintf("%d-%d", coord.X, coord.Y)
		// because of how we initialize, we can end up where origin is, if so, skip looking here
		if coord.X == sb.Home.X && coord.Y == sb.Home.Y {
			continue
		}
		if string(plane[coord.X][coord.Y]) == charac {
			if coordinatesSeen[coordString] == 0 {
				count += 1
			}
		}
		coordinatesSeen[coordString] = coordinatesSeen[coordString] + 1
	}
	return count
}

func (c *Coordinate) FindCharacter(charac string, plane Plane2, vert Vertical, horz Horizontal) {
	if string(plane[c.X][c.Y]) != charac {
		c.identified = true
	} else {
		// identify new coordinates
		var newX, newY int
		switch horz {
		case Left:
			newY = c.Y - 1
		case Forward:
			newY = c.Y
		case Right:
			newY = c.Y + 1
		}
		switch vert {
		case Up:
			newX = c.X - 1
		case Straight:
			newX = c.X
		case Down:
			newX = c.X + 1
		}
		// check the new coordinate and move
		if newX < 0 || newX > c.MaxIndexX || newY < 0 || newY > c.MaxIndexY {
			// we've moved somewhere that is outside the bounds, so we can't look any further
			c.identified = true
		} else {
			// the move is valid and we should look again to see if we move or have found the right place
			c.X = newX
			c.Y = newY
			c.FindCharacter(charac, plane, vert, horz)
		}
	}
}

func CalcSeatChanges(plane Plane2) [][2]int {
	changes := [][2]int{}
	for x := 0; x < len(plane); x++ {
		for y := 0; y < len(plane[0]); y++ {
			home := Coordinate{X: x, Y: y, MaxIndexX: len(plane) - 1, MaxIndexY: len(plane[0]) - 1}
			sb := SearchBox{Home: home}
			sb.FindBoundaries(".", plane)
			value := string(plane[x][y])
			if value == "." {
				// coordinate is floor, skip
				continue
			} else if (value == "#" && sb.CountCharInBoundingBox("#", plane) >= 5) || (value == "L" && sb.CountCharInBoundingBox("#", plane) == 0) {
				changes = append(changes, [2]int{x, y})
			} else {
				// not needing a change
				continue
			}
		}
	}
	return changes
}

func BuildPlane(plane Plane2) Plane2 {
	constructor := []string{}
	seatChanges := CalcSeatChanges(plane)
	for rowIdx, row := range plane {
		newRow := ""
		for colIdx, val := range row {
			changedVal := false
			for _, change := range seatChanges {
				if change == [2]int{rowIdx, colIdx} && string(val) == "#" {
					newRow += "L"
					changedVal = true
				} else if change == [2]int{rowIdx, colIdx} && string(val) == "L" {
					newRow += "#"
					changedVal = true
				}
			}
			if !changedVal {
				newRow += string(val)
			}
		}
		constructor = append(constructor, newRow)
	}
	return Plane2(constructor)
}
