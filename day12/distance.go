package distance

import (
	"math"
	"strconv"
)

type Direction string

const (
	East  Direction = "E"
	West            = "W"
	North           = "N"
	South           = "S"
)

type Ferry struct {
	East          int
	North         int
	Facing        Direction
	Degrees       int
	WaypointEast  int
	WaypointNorth int
}

func CreateFerry() Ferry {
	return Ferry{Facing: East, Degrees: 90, WaypointEast: 10, WaypointNorth: 1}
}

func (f *Ferry) MoveForward(units int) {
	switch f.Facing {
	case North:
		f.North += units
	case South:
		f.North -= units
	case East:
		f.East += units
	case West:
		f.East -= units
	}
}

func (f *Ferry) findFacingByDegrees() {
	if f.Degrees == 360 {
		f.Degrees = 0
	}
	if f.Degrees >= 0 && f.Degrees < 90 {
		f.Facing = North
	} else if f.Degrees >= 90 && f.Degrees < 180 {
		f.Facing = East
	} else if f.Degrees >= 180 && f.Degrees < 270 {
		f.Facing = South
	} else {
		f.Facing = West
	}
}

func (f *Ferry) TurnLeft(degrees int) {
	newDegrees := f.Degrees - degrees
	if newDegrees < 0 {
		newDegrees += 360
	}
	f.Degrees = newDegrees
	f.findFacingByDegrees()
}

func (f *Ferry) TurnRight(degrees int) {
	newDegrees := f.Degrees + degrees
	if newDegrees > 360 {
		newDegrees -= 360
	}
	f.Degrees = newDegrees
	f.findFacingByDegrees()
}

func (f *Ferry) waypointRightNorth() int {
	return f.WaypointEast * -1
}

func (f *Ferry) waypointLeftNorth() int {
	return f.WaypointEast
}

func (f *Ferry) waypointRightEast() int {
	return f.WaypointNorth
}

func (f *Ferry) waypointLeftEast() int {
	return f.WaypointNorth * -1
}

func (f *Ferry) RotateWaypointRight(degrees int) {
	numberRotations := degrees / 90
	for i := 0; i < numberRotations; i++ {
		newNorth := f.waypointRightNorth()
		newEast := f.waypointRightEast()
		f.WaypointNorth = newNorth
		f.WaypointEast = newEast
	}
}

func (f *Ferry) RotateWaypointLeft(degrees int) {
	numberRotations := degrees / 90
	for i := 0; i < numberRotations; i++ {
		newNorth := f.waypointLeftNorth()
		newEast := f.waypointLeftEast()
		f.WaypointNorth = newNorth
		f.WaypointEast = newEast
	}
}

func (f *Ferry) MoveByDirection(direction Direction, units int) {
	switch direction {
	case North:
		f.North += units
	case South:
		f.North -= units
	case East:
		f.East += units
	case West:
		f.East -= units
	}
}

func (f *Ferry) MoveWaypointByDirection(direction Direction, units int) {
	switch direction {
	case North:
		f.WaypointNorth += units
	case South:
		f.WaypointNorth -= units
	case East:
		f.WaypointEast += units
	case West:
		f.WaypointEast -= units
	}
}

func (f *Ferry) MoveTowardsWaypoint(units int) {
	f.North += units * f.WaypointNorth
	f.East += units * f.WaypointEast
}

func (f *Ferry) CurrentManhattanDistance() int {
	return int(math.Abs(float64(f.East)) + math.Abs(float64(f.North)))
}

func (f *Ferry) ProcessDirection(direction string, useWaypoint bool) {
	direc := string(direction[0])
	value, _ := strconv.Atoi(direction[1:])
	switch useWaypoint {
	case false:
		switch direc {
		case "N":
			f.MoveByDirection(North, value)
		case "E":
			f.MoveByDirection(East, value)
		case "S":
			f.MoveByDirection(South, value)
		case "W":
			f.MoveByDirection(West, value)
		case "L":
			f.TurnLeft(value)
		case "R":
			f.TurnRight(value)
		case "F":
			f.MoveForward(value)
		}
	case true:
		switch direc {
		case "N":
			f.MoveWaypointByDirection(North, value)
		case "E":
			f.MoveWaypointByDirection(East, value)
		case "S":
			f.MoveWaypointByDirection(South, value)
		case "W":
			f.MoveWaypointByDirection(West, value)
		case "L":
			f.RotateWaypointLeft(value)
		case "R":
			f.RotateWaypointRight(value)
		case "F":
			f.MoveTowardsWaypoint(value)
		}
	}
}
