package distance

import (
	"testing"
)

func testMoving(expected int, actual int) func(*testing.T) {
	return func(t *testing.T) {
		if expected != actual {
			t.Errorf("Should have moved %d units, moved %d units instead", expected, actual)
		}
	}
}

func TestMoveForward(t *testing.T) {
	ferry := Ferry{Facing: East}
	ferry.MoveForward(10)
	t.Run("Move east", testMoving(10, ferry.East))

	ferry.Facing = North
	ferry.MoveForward(10)
	t.Run("Move north", testMoving(10, ferry.North))

	ferry.Facing = West
	ferry.MoveForward(10)
	t.Run("Move west", testMoving(0, ferry.East))

	ferry.Facing = South
	ferry.MoveForward(10)
	t.Run("Move south", testMoving(0, ferry.North))
}

func TestMoveByDirection(t *testing.T) {
	ferry := Ferry{Facing: East}
	ferry.MoveByDirection(North, 10)
	t.Run("Move by direction North", testMoving(10, ferry.North))

	ferry.MoveByDirection(South, 10)
	t.Run("Move by direction South", testMoving(0, ferry.North))

	ferry.MoveByDirection(East, 10)
	t.Run("Move by direction East", testMoving(10, ferry.East))

	ferry.MoveByDirection(West, 10)
	t.Run("Move by direction West", testMoving(0, ferry.East))
}

func TestMoveWaypointByDirection(t *testing.T) {
	ferry := CreateFerry()
	ferry.MoveWaypointByDirection(North, 10)
	t.Run("Move waypoint by 10 units North", testMoving(11, ferry.WaypointNorth))

	ferry.MoveWaypointByDirection(South, 10)
	t.Run("Move waypoint by 10 units South", testMoving(1, ferry.WaypointNorth))

	ferry.MoveWaypointByDirection(East, 10)
	t.Run("Move waypoint by 10 units East", testMoving(20, ferry.WaypointEast))

	ferry.MoveWaypointByDirection(West, 10)
	t.Run("Move waypoint by 10 units West", testMoving(10, ferry.WaypointEast))
}

func TestMoveTowardsWaypoint(t *testing.T) {
	ferry := CreateFerry()
	ferry.MoveTowardsWaypoint(10)
	t.Run("Move towards waypoint with value of 10, North", testMoving(10, ferry.North))
	t.Run("Move towards waypoint with value of 10, East", testMoving(100, ferry.East))
}

func TestRotateWaypoint(t *testing.T) {
	ferry := CreateFerry()

	ferry.RotateWaypointRight(90)
	t.Run("Right by 90, North", testMoving(-10, ferry.WaypointNorth))
	t.Run("Right by 90, East", testMoving(1, ferry.WaypointEast))

	ferry.RotateWaypointRight(180)
	t.Run("Right by 180, North", testMoving(10, ferry.WaypointNorth))
	t.Run("Right by 180, East", testMoving(-1, ferry.WaypointEast))

	ferry.RotateWaypointRight(270)
	t.Run("Right by 180, North", testMoving(-1, ferry.WaypointNorth))
	t.Run("Right by 180, East", testMoving(-10, ferry.WaypointEast))

	ferry.RotateWaypointLeft(90)
	t.Run("Left by 90, North", testMoving(-10, ferry.WaypointNorth))
	t.Run("Left by 90, East", testMoving(1, ferry.WaypointEast))

	ferry.RotateWaypointLeft(180)
	t.Run("Left by 180, North", testMoving(10, ferry.WaypointNorth))
	t.Run("Left by 180, East", testMoving(-1, ferry.WaypointEast))

	ferry.RotateWaypointLeft(270)
	t.Run("Left by 270, North", testMoving(1, ferry.WaypointNorth))
	t.Run("Left by 270, East", testMoving(10, ferry.WaypointEast))
}

func testFacing(expected Direction, actual Direction) func(*testing.T) {
	return func(t *testing.T) {
		if expected != actual {
			t.Errorf("Should be facing %s but facing %s instead", expected, actual)
		}
	}
}

func TestChangeDirection(t *testing.T) {
	ferry := CreateFerry()
	ferry.TurnLeft(90)
	t.Run("Turn left 90 degrees", testFacing(North, ferry.Facing))

	ferry.TurnLeft(180)
	t.Run("Turn left 180 degrees", testFacing(South, ferry.Facing))

	ferry.TurnLeft(270)
	t.Run("Turn left 270 degrees", testFacing(West, ferry.Facing))

	ferry.TurnRight(90)
	t.Run("Turn right 90 degrees", testFacing(North, ferry.Facing))

	ferry.TurnRight(180)
	t.Run("Turn right 180 degrees", testFacing(South, ferry.Facing))

	ferry.TurnRight(270)
	t.Run("Turn right 270 degrees", testFacing(East, ferry.Facing))
}

func TestManhattanDirection(t *testing.T) {
	ferry := Ferry{East: 17, North: -8}
	resp := ferry.CurrentManhattanDistance()
	if resp != 25 {
		t.Errorf("Manhattan distance should be 25")
	}
}

func TestFollowDirections(t *testing.T) {
	ferry := CreateFerry()
	directions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	for _, direction := range directions {
		ferry.ProcessDirection(direction, false)
	}
	if ferry.CurrentManhattanDistance() != 25 {
		t.Errorf("Should be manhattan distance of 25, ferry is currently %+v", ferry)
	}
}

func TestWaypointDirections(t *testing.T) {
	ferry := CreateFerry()
	directions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	for _, direction := range directions {
		ferry.ProcessDirection(direction, true)
	}
	if ferry.CurrentManhattanDistance() != 286 {
		t.Errorf("Should be manhattan distance of 286, ferry is currently %+v", ferry)
	}
}
