package coordinate

import (
	"fmt"
)

// Direction represents a direction
type Direction interface {
	// Opposite returns the opposite direction
	Opposite() Direction
	// ToCoordinate2D returns the coordinate2D representation of the direction
	ToCoordinate2D() Coordinate2D[int]
	// toInt converts the direction to an integer. It is used internally to
	// calculate the next direction easier. It also 'closes' the interface
	toInt() int
	// TurnRight returns the direction after turning 90 degrees to the right
	TurnRight() Direction
	// TurnRight45 returns the direction after turning 45 degrees to the right
	TurnRight45() Direction
	// TurnLeft returns the direction after turning 90 degrees to the left
	TurnLeft() Direction
	// TurnLeft45 returns the direction after turning 45 degrees to the left
	TurnLeft45() Direction
	fmt.Stringer
}

const (
	up = iota
	upRight
	right
	downRight
	down
	downLeft
	left
	upLeft
	DirectionCount
)

// coordinates stores the coordinates for each direction
var coordinates = []Coordinate2D[int]{
	{X: 0, Y: -1},  // up
	{X: 1, Y: -1},  // up-right
	{X: 1, Y: 0},   // right
	{X: 1, Y: 1},   // down-right
	{X: 0, Y: 1},   // down
	{X: -1, Y: 1},  // down-left
	{X: -1, Y: 0},  // left
	{X: -1, Y: -1}, // up-left
}

// Format stores the string representation of each direction. It is used for parsing and printing.
// The order is important and should match the order of the directions going clockwise starting from up.
// There are a couple predefined formats that can be used.
var Format = Default

var (
	// Default stores the default string representation of each direction
	Default = [DirectionCount]string{"up", "up-right", "right", "down-right", "down", "down-left", "left", "up-left"}
	// DefaultCamel stores the default string representation of each direction in camel case
	DefaultCamel = [DirectionCount]string{"Up", "UpRight", "Right", "DownRight", "Down", "DownLeft", "Left", "UpLeft"}
	// Short stores the short string representation of each direction
	Short = [DirectionCount]string{"u", "ur", "r", "dr", "d", "dl", "l", "ul"}
	// ShortUpper stores the short string representation of each direction in upper case
	ShortUpper = [DirectionCount]string{"U", "UR", "R", "DR", "D", "DL", "L", "UL"}
	// Compass stores the compass string representation of each direction
	Compass = [DirectionCount]string{"north", "north-east", "east", "south-east", "south", "south-west", "west", "north-west"}
	// CompassCamel stores the full compass string representation of each direction in camel case
	CompassCamel = [DirectionCount]string{"North", "NorthEast", "East", "SouthEast", "South", "SouthWest", "West", "NorthWest"}
	// CompassShort stores the short compass string representation of each direction
	CompassShort = [DirectionCount]string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}
	// CompassShortUpper stores the short compass string representation of each direction in upper case
	CompassShortUpper = [DirectionCount]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	// Characters stores the characters for each direction. It only contains characters for the 4 main directions.
	Characters = [DirectionCount]string{"^", "", ">", "", "v", "", "<", ""}
)

// Up represents the up direction
func Up() Direction {
	return fromInt(up)
}

// UpRight represents the up-right direction
func UpRight() Direction {
	return fromInt(upRight)
}

// Right represents the right direction
func Right() Direction {
	return fromInt(right)
}

// DownRight represents the down-right direction
func DownRight() Direction {
	return fromInt(downRight)
}

// Down represents the down direction
func Down() Direction {
	return fromInt(down)
}

// DownLeft represents the down-left direction
func DownLeft() Direction {
	return fromInt(downLeft)
}

// Left represents the left direction
func Left() Direction {
	return fromInt(left)
}

// UpLeft represents the up-left direction
func UpLeft() Direction {
	return fromInt(upLeft)
}

// Parse parses a string to a direction based on Format
func Parse(s string) (Direction, error) {
	for i, f := range Format {
		if f == s {
			return fromInt(i), nil
		}
	}
	return nil, fmt.Errorf("unknown direction: %s", s)
}

// fromInt converts an integer to a direction based on predefined constants
func fromInt(i int) Direction {
	return direction(i)
}

// direction implements the Direction interface
type direction int

// Opposite returns the opposite direction
func (d direction) Opposite() Direction {
	return fromInt((d.toInt() + DirectionCount/2) % DirectionCount)
}

// ToCoordinate2D returns the coordinate2D representation of the direction
func (d direction) ToCoordinate2D() Coordinate2D[int] {
	return coordinates[d.toInt()]
}

// toInt converts the direction to an integer
func (d direction) toInt() int {
	return int(d)
}

// TurnRight returns the direction after turning 90 degrees to the right
func (d direction) TurnRight() Direction {
	return fromInt((d.toInt() + DirectionCount/4) % DirectionCount)
}

// TurnRight45 returns the direction after turning 45 degrees to the right
func (d direction) TurnRight45() Direction {
	return fromInt((d.toInt() + DirectionCount/8) % DirectionCount)
}

// TurnLeft returns the direction after turning 90 degrees to the left
func (d direction) TurnLeft() Direction {
	return fromInt((d.toInt() + DirectionCount*3/4) % DirectionCount)
}

// TurnLeft45 returns the direction after turning 45 degrees to the left
func (d direction) TurnLeft45() Direction {
	return fromInt((d.toInt() + DirectionCount*7/8) % DirectionCount)
}

// String returns the string representation of the direction
func (d direction) String() string {
	return Format[d.toInt()]
}
