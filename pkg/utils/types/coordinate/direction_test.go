package coordinate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUp(t *testing.T) {
	assert.Equal(t, Up(), fromInt(up))
}

func TestUpRight(t *testing.T) {
	assert.Equal(t, UpRight(), fromInt(upRight))
}

func TestRight(t *testing.T) {
	assert.Equal(t, Right(), fromInt(right))
}

func TestDownRight(t *testing.T) {
	assert.Equal(t, DownRight(), fromInt(downRight))
}

func TestDown(t *testing.T) {
	assert.Equal(t, Down(), fromInt(down))
}

func TestDownLeft(t *testing.T) {
	assert.Equal(t, DownLeft(), fromInt(downLeft))
}

func TestLeft(t *testing.T) {
	assert.Equal(t, Left(), fromInt(left))
}

func TestUpLeft(t *testing.T) {
	assert.Equal(t, UpLeft(), fromInt(upLeft))
}

func TestDiagonals(t *testing.T) {
	assert.Equal(t, []Direction{UpRight(), DownRight(), DownLeft(), UpLeft()}, Diagonals())
}

func TestHorizontals(t *testing.T) {
	assert.Equal(t, []Direction{Right(), Left()}, Horizontals())
}

func TestStraights(t *testing.T) {
	assert.Equal(t, []Direction{Up(), Right(), Down(), Left()}, Straights())
}

func TestVerticals(t *testing.T) {
	assert.Equal(t, []Direction{Up(), Down()}, Verticals())
}

func TestParse(t *testing.T) {
	type testCase struct {
		testName      string
		format        [DirectionCount]string
		s             string
		expected      Direction
		expectedError bool
	}
	customFormat := [DirectionCount]string{"1", "2", "3", "4", "5", "6", "7", "8"}
	testCases := []testCase{
		{"Default up", Default, "up", Up(), false},
		{"Default up-right", Default, "up-right", UpRight(), false},
		{"Default right", Default, "right", Right(), false},
		{"Default down-right", Default, "down-right", DownRight(), false},
		{"Default down", Default, "down", Down(), false},
		{"Default down-left", Default, "down-left", DownLeft(), false},
		{"Default left", Default, "left", Left(), false},
		{"Default up-left", Default, "up-left", UpLeft(), false},
		{"Default invalid", Default, "invalid", nil, true},
		{"DefaultCamel up", DefaultCamel, "Up", Up(), false},
		{"DefaultCamel up-right", DefaultCamel, "UpRight", UpRight(), false},
		{"DefaultCamel right", DefaultCamel, "Right", Right(), false},
		{"DefaultCamel down-right", DefaultCamel, "DownRight", DownRight(), false},
		{"DefaultCamel down", DefaultCamel, "Down", Down(), false},
		{"DefaultCamel down-left", DefaultCamel, "DownLeft", DownLeft(), false},
		{"DefaultCamel left", DefaultCamel, "Left", Left(), false},
		{"DefaultCamel up-left", DefaultCamel, "UpLeft", UpLeft(), false},
		{"DefaultCamel invalid", DefaultCamel, "Invalid", nil, true},
		{"Short up", Short, "u", Up(), false},
		{"Short up-right", Short, "ur", UpRight(), false},
		{"Short right", Short, "r", Right(), false},
		{"Short down-right", Short, "dr", DownRight(), false},
		{"Short down", Short, "d", Down(), false},
		{"Short down-left", Short, "dl", DownLeft(), false},
		{"Short left", Short, "l", Left(), false},
		{"Short up-left", Short, "ul", UpLeft(), false},
		{"Short invalid", Short, "i", nil, true},
		{"ShortUpper up", ShortUpper, "U", Up(), false},
		{"ShortUpper up-right", ShortUpper, "UR", UpRight(), false},
		{"ShortUpper right", ShortUpper, "R", Right(), false},
		{"ShortUpper down-right", ShortUpper, "DR", DownRight(), false},
		{"ShortUpper down", ShortUpper, "D", Down(), false},
		{"ShortUpper down-left", ShortUpper, "DL", DownLeft(), false},
		{"ShortUpper left", ShortUpper, "L", Left(), false},
		{"ShortUpper up-left", ShortUpper, "UL", UpLeft(), false},
		{"ShortUpper invalid", ShortUpper, "I", nil, true},
		{"Compass north", Compass, "north", Up(), false},
		{"Compass north-east", Compass, "north-east", UpRight(), false},
		{"Compass east", Compass, "east", Right(), false},
		{"Compass south-east", Compass, "south-east", DownRight(), false},
		{"Compass south", Compass, "south", Down(), false},
		{"Compass south-west", Compass, "south-west", DownLeft(), false},
		{"Compass west", Compass, "west", Left(), false},
		{"Compass north-west", Compass, "north-west", UpLeft(), false},
		{"Compass invalid", Compass, "invalid", nil, true},
		{"CompassCamel north", CompassCamel, "North", Up(), false},
		{"CompassCamel north-east", CompassCamel, "NorthEast", UpRight(), false},
		{"CompassCamel east", CompassCamel, "East", Right(), false},
		{"CompassCamel south-east", CompassCamel, "SouthEast", DownRight(), false},
		{"CompassCamel south", CompassCamel, "South", Down(), false},
		{"CompassCamel south-west", CompassCamel, "SouthWest", DownLeft(), false},
		{"CompassCamel west", CompassCamel, "West", Left(), false},
		{"CompassCamel north-west", CompassCamel, "NorthWest", UpLeft(), false},
		{"CompassCamel invalid", CompassCamel, "Invalid", nil, true},
		{"CompassShort north", CompassShort, "n", Up(), false},
		{"CompassShort north-east", CompassShort, "ne", UpRight(), false},
		{"CompassShort east", CompassShort, "e", Right(), false},
		{"CompassShort south-east", CompassShort, "se", DownRight(), false},
		{"CompassShort south", CompassShort, "s", Down(), false},
		{"CompassShort south-west", CompassShort, "sw", DownLeft(), false},
		{"CompassShort west", CompassShort, "w", Left(), false},
		{"CompassShort north-west", CompassShort, "nw", UpLeft(), false},
		{"CompassShort invalid", CompassShort, "i", nil, true},
		{"CompassShortUpper north", CompassShortUpper, "N", Up(), false},
		{"CompassShortUpper north-east", CompassShortUpper, "NE", UpRight(), false},
		{"CompassShortUpper east", CompassShortUpper, "E", Right(), false},
		{"CompassShortUpper south-east", CompassShortUpper, "SE", DownRight(), false},
		{"CompassShortUpper south", CompassShortUpper, "S", Down(), false},
		{"CompassShortUpper south-west", CompassShortUpper, "SW", DownLeft(), false},
		{"CompassShortUpper west", CompassShortUpper, "W", Left(), false},
		{"CompassShortUpper north-west", CompassShortUpper, "NW", UpLeft(), false},
		{"CompassShortUpper invalid", CompassShortUpper, "I", nil, true},
		{"Characters up", Characters, "^", Up(), false},
		{"Characters right", Characters, ">", Right(), false},
		{"Characters down", Characters, "v", Down(), false},
		{"Characters left", Characters, "<", Left(), false},
		{"Characters invalid", Characters, "i", nil, true},
		{"Custom up", customFormat, "1", Up(), false},
		{"Custom up-right", customFormat, "2", UpRight(), false},
		{"Custom right", customFormat, "3", Right(), false},
		{"Custom down-right", customFormat, "4", DownRight(), false},
		{"Custom down", customFormat, "5", Down(), false},
		{"Custom down-left", customFormat, "6", DownLeft(), false},
		{"Custom left", customFormat, "7", Left(), false},
		{"Custom up-left", customFormat, "8", UpLeft(), false},
		{"Custom invalid", customFormat, "9", nil, true},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			Format = tc.format
			result, err := Parse(tc.s)
			if tc.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}

func TestFromInt(t *testing.T) {
	testCases := []struct {
		testName string
		i        int
		expected Direction
	}{
		{"Up", up, Up()},
		{"UpRight", upRight, UpRight()},
		{"Right", right, Right()},
		{"DownRight", downRight, DownRight()},
		{"Down", down, Down()},
		{"DownLeft", downLeft, DownLeft()},
		{"Left", left, Left()},
		{"UpLeft", upLeft, UpLeft()},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, fromInt(tc.i))
		})
	}
}

func TestDiagonal(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected bool
	}{
		{"Up", Up(), false},
		{"UpRight", UpRight(), true},
		{"Right", Right(), false},
		{"DownRight", DownRight(), true},
		{"Down", Down(), false},
		{"DownLeft", DownLeft(), true},
		{"Left", Left(), false},
		{"UpLeft", UpLeft(), true},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.Diagonal())
		})
	}
}

func TestHorizontal(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected bool
	}{
		{"Up", Up(), false},
		{"UpRight", UpRight(), false},
		{"Right", Right(), true},
		{"DownRight", DownRight(), false},
		{"Down", Down(), false},
		{"DownLeft", DownLeft(), false},
		{"Left", Left(), true},
		{"UpLeft", UpLeft(), false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.Horizontal())
		})
	}
}

func TestOpposite(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected Direction
	}{
		{"Up", Up(), Down()},
		{"UpRight", UpRight(), DownLeft()},
		{"Right", Right(), Left()},
		{"DownRight", DownRight(), UpLeft()},
		{"Down", Down(), Up()},
		{"DownLeft", DownLeft(), UpRight()},
		{"Left", Left(), Right()},
		{"UpLeft", UpLeft(), DownRight()},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.Opposite())
		})
	}
}

func TestToCoordinate2D(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected Coordinate2D[int]
	}{
		{"Up", Up(), Coordinate2D[int]{X: 0, Y: -1}},
		{"UpRight", UpRight(), Coordinate2D[int]{X: 1, Y: -1}},
		{"Right", Right(), Coordinate2D[int]{X: 1, Y: 0}},
		{"DownRight", DownRight(), Coordinate2D[int]{X: 1, Y: 1}},
		{"Down", Down(), Coordinate2D[int]{X: 0, Y: 1}},
		{"DownLeft", DownLeft(), Coordinate2D[int]{X: -1, Y: 1}},
		{"Left", Left(), Coordinate2D[int]{X: -1, Y: 0}},
		{"UpLeft", UpLeft(), Coordinate2D[int]{X: -1, Y: -1}},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.ToCoordinate2D())
		})
	}
}

func TestToInt(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected int
	}{
		{"Up", Up(), up},
		{"UpRight", UpRight(), upRight},
		{"Right", Right(), right},
		{"DownRight", DownRight(), downRight},
		{"Down", Down(), down},
		{"DownLeft", DownLeft(), downLeft},
		{"Left", Left(), left},
		{"UpLeft", UpLeft(), upLeft},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.toInt())
		})
	}
}

func TestTurnRight(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected Direction
	}{
		{"Up", Up(), Right()},
		{"UpRight", UpRight(), DownRight()},
		{"Right", Right(), Down()},
		{"DownRight", DownRight(), DownLeft()},
		{"Down", Down(), Left()},
		{"DownLeft", DownLeft(), UpLeft()},
		{"Left", Left(), Up()},
		{"UpLeft", UpLeft(), UpRight()},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.TurnRight())
		})
	}
}

func TestTurnRight45(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected Direction
	}{
		{"Up", Up(), UpRight()},
		{"UpRight", UpRight(), Right()},
		{"Right", Right(), DownRight()},
		{"DownRight", DownRight(), Down()},
		{"Down", Down(), DownLeft()},
		{"DownLeft", DownLeft(), Left()},
		{"Left", Left(), UpLeft()},
		{"UpLeft", UpLeft(), Up()},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.TurnRight45())
		})
	}
}

func TestTurnLeft(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected Direction
	}{
		{"Up", Up(), Left()},
		{"UpRight", UpRight(), UpLeft()},
		{"Right", Right(), Up()},
		{"DownRight", DownRight(), UpRight()},
		{"Down", Down(), Right()},
		{"DownLeft", DownLeft(), DownRight()},
		{"Left", Left(), Down()},
		{"UpLeft", UpLeft(), DownLeft()},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.TurnLeft())
		})
	}
}

func TestTurnLeft45(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected Direction
	}{
		{"Up", Up(), UpLeft()},
		{"UpRight", UpRight(), Up()},
		{"Right", Right(), UpRight()},
		{"DownRight", DownRight(), Right()},
		{"Down", Down(), DownRight()},
		{"DownLeft", DownLeft(), Down()},
		{"Left", Left(), DownLeft()},
		{"UpLeft", UpLeft(), Left()},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.TurnLeft45())
		})
	}
}

func TestString(t *testing.T) {
	customFormat := [DirectionCount]string{"1", "2", "3", "4", "5", "6", "7", "8"}
	testCases := []struct {
		testName string
		d        Direction
		format   [DirectionCount]string
		expected string
	}{
		{"Default up", Up(), Default, "up"},
		{"Default up-right", UpRight(), Default, "up-right"},
		{"Default right", Right(), Default, "right"},
		{"Default down-right", DownRight(), Default, "down-right"},
		{"Default down", Down(), Default, "down"},
		{"Default down-left", DownLeft(), Default, "down-left"},
		{"Default left", Left(), Default, "left"},
		{"Default up-left", UpLeft(), Default, "up-left"},
		{"DefaultCamel up", Up(), DefaultCamel, "Up"},
		{"DefaultCamel up-right", UpRight(), DefaultCamel, "UpRight"},
		{"DefaultCamel right", Right(), DefaultCamel, "Right"},
		{"DefaultCamel down-right", DownRight(), DefaultCamel, "DownRight"},
		{"DefaultCamel down", Down(), DefaultCamel, "Down"},
		{"DefaultCamel down-left", DownLeft(), DefaultCamel, "DownLeft"},
		{"DefaultCamel left", Left(), DefaultCamel, "Left"},
		{"DefaultCamel up-left", UpLeft(), DefaultCamel, "UpLeft"},
		{"Short up", Up(), Short, "u"},
		{"Short up-right", UpRight(), Short, "ur"},
		{"Short right", Right(), Short, "r"},
		{"Short down-right", DownRight(), Short, "dr"},
		{"Short down", Down(), Short, "d"},
		{"Short down-left", DownLeft(), Short, "dl"},
		{"Short left", Left(), Short, "l"},
		{"Short up-left", UpLeft(), Short, "ul"},
		{"ShortUpper up", Up(), ShortUpper, "U"},
		{"ShortUpper up-right", UpRight(), ShortUpper, "UR"},
		{"ShortUpper right", Right(), ShortUpper, "R"},
		{"ShortUpper down-right", DownRight(), ShortUpper, "DR"},
		{"ShortUpper down", Down(), ShortUpper, "D"},
		{"ShortUpper down-left", DownLeft(), ShortUpper, "DL"},
		{"ShortUpper left", Left(), ShortUpper, "L"},
		{"ShortUpper up-left", UpLeft(), ShortUpper, "UL"},
		{"Compass north", Up(), Compass, "north"},
		{"Compass north-east", UpRight(), Compass, "north-east"},
		{"Compass east", Right(), Compass, "east"},
		{"Compass south-east", DownRight(), Compass, "south-east"},
		{"Compass south", Down(), Compass, "south"},
		{"Compass south-west", DownLeft(), Compass, "south-west"},
		{"Compass west", Left(), Compass, "west"},
		{"Compass north-west", UpLeft(), Compass, "north-west"},
		{"CompassCamel north", Up(), CompassCamel, "North"},
		{"CompassCamel north-east", UpRight(), CompassCamel, "NorthEast"},
		{"CompassCamel east", Right(), CompassCamel, "East"},
		{"CompassCamel south-east", DownRight(), CompassCamel, "SouthEast"},
		{"CompassCamel south", Down(), CompassCamel, "South"},
		{"CompassCamel south-west", DownLeft(), CompassCamel, "SouthWest"},
		{"CompassCamel west", Left(), CompassCamel, "West"},
		{"CompassCamel north-west", UpLeft(), CompassCamel, "NorthWest"},
		{"CompassShort north", Up(), CompassShort, "n"},
		{"CompassShort north-east", UpRight(), CompassShort, "ne"},
		{"CompassShort east", Right(), CompassShort, "e"},
		{"CompassShort south-east", DownRight(), CompassShort, "se"},
		{"CompassShort south", Down(), CompassShort, "s"},
		{"CompassShort south-west", DownLeft(), CompassShort, "sw"},
		{"CompassShort west", Left(), CompassShort, "w"},
		{"CompassShort north-west", UpLeft(), CompassShort, "nw"},
		{"CompassShortUpper north", Up(), CompassShortUpper, "N"},
		{"CompassShortUpper north-east", UpRight(), CompassShortUpper, "NE"},
		{"CompassShortUpper east", Right(), CompassShortUpper, "E"},
		{"CompassShortUpper south-east", DownRight(), CompassShortUpper, "SE"},
		{"CompassShortUpper south", Down(), CompassShortUpper, "S"},
		{"CompassShortUpper south-west", DownLeft(), CompassShortUpper, "SW"},
		{"CompassShortUpper west", Left(), CompassShortUpper, "W"},
		{"CompassShortUpper north-west", UpLeft(), CompassShortUpper, "NW"},
		{"Characters up", Up(), Characters, "^"},
		{"Characters right", Right(), Characters, ">"},
		{"Characters down", Down(), Characters, "v"},
		{"Characters left", Left(), Characters, "<"},
		{"Custom up", Up(), customFormat, "1"},
		{"Custom up-right", UpRight(), customFormat, "2"},
		{"Custom right", Right(), customFormat, "3"},
		{"Custom down-right", DownRight(), customFormat, "4"},
		{"Custom down", Down(), customFormat, "5"},
		{"Custom down-left", DownLeft(), customFormat, "6"},
		{"Custom left", Left(), customFormat, "7"},
		{"Custom up-left", UpLeft(), customFormat, "8"},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			Format = tc.format
			assert.Equal(t, tc.expected, tc.d.String())
		})
	}
}

func TestVertical(t *testing.T) {
	testCases := []struct {
		testName string
		d        Direction
		expected bool
	}{
		{"Up", Up(), true},
		{"UpRight", UpRight(), false},
		{"Right", Right(), false},
		{"DownRight", DownRight(), false},
		{"Down", Down(), true},
		{"DownLeft", DownLeft(), false},
		{"Left", Left(), false},
		{"UpLeft", UpLeft(), false},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.d.Vertical())
		})
	}
}
