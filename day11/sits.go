package sits

type Plane []string

type SurroundingBox struct {
	X    int
	Y    int
	XMin int
	XMax int
	YMin int
	YMax int
}

func (sb *SurroundingBox) FindSurroundingBox(maxX int, maxY int) {
	var XMin, XMax, YMin, YMax int
	if sb.X == 0 {
		// point is on the left edge of the plane
		XMin = 0
	} else {
		XMin = sb.X - 1
	}
	if sb.X == maxX {
		// point is on the right edge of the plane
		XMax = maxX
	} else {
		XMax = sb.X + 1
	}
	if sb.Y == 0 {
		// point is on the top line of the plane
		YMin = 0
	} else {
		YMin = sb.Y - 1
	}
	if sb.Y == maxY {
		// point is on the bottom line of the plane
		YMax = maxY
	} else {
		YMax = sb.Y + 1
	}
	sb.XMin = XMin
	sb.XMax = XMax
	sb.YMin = YMin
	sb.YMax = YMax
}

func CountCharInBoundingBox(charac string, plane Plane, box SurroundingBox) int {
	count := 0
	for x := box.XMin; x < box.XMax+1; x++ {
		for y := box.YMin; y < box.YMax+1; y++ {
			if x == box.X && y == box.Y {
				continue
			}
			if string(plane[x][y]) == charac {
				count += 1
			}
		}
	}
	return count
}

func CalcSeatChanges(plane Plane) [][2]int {
	maxX := len(plane) - 1
	maxY := len(plane[0]) - 1
	changes := [][2]int{}
	for x := 0; x < len(plane); x++ {
		for y := 0; y < len(plane[0]); y++ {
			box := SurroundingBox{X: x, Y: y}
			box.FindSurroundingBox(maxX, maxY)
			value := string(plane[x][y])
			if value == "." {
				// coordinate is floor, skip
				continue
			} else if (value == "#" && CountCharInBoundingBox("#", plane, box) >= 4) || (value == "L" && CountCharInBoundingBox("#", plane, box) == 0) {
				changes = append(changes, [2]int{x, y})
			} else {
				// not needing a change
				continue
			}
		}
	}
	return changes
}

func BuildPlane(plane Plane) Plane {
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
	return Plane(constructor)
}
