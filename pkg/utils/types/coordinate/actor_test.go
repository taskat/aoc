package coordinate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewActor(t *testing.T) {
	testCases := []struct {
		testName  string
		position  Coordinate2D[int]
		direction Direction
	}{
		{
			testName:  "Actor at origin facing up",
			position:  Coordinate2D[int]{X: 0, Y: 0},
			direction: Up(),
		},
		{
			testName:  "Actor at (2,3) facing right",
			position:  Coordinate2D[int]{X: 2, Y: 3},
			direction: Right(),
		},
		{
			testName:  "Actor at (-1,-1) facing down-left",
			position:  Coordinate2D[int]{X: -1, Y: -1},
			direction: DownLeft(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actor := NewActor(tc.position, tc.direction)
			assert.Equal(t, tc.position, actor.Position)
			assert.Equal(t, tc.direction, actor.Direction)
		})
	}
}

// TestActorMove tests the Move method of the Actor struct
func TestActorMove(t *testing.T) {
	testCases := []struct {
		testName       string
		initialPos     Coordinate2D[int]
		direction      Direction
		expectedNewPos Coordinate2D[int]
	}{
		{
			testName:       "Move up from origin",
			initialPos:     Coordinate2D[int]{X: 0, Y: 0},
			direction:      Up(),
			expectedNewPos: Coordinate2D[int]{X: 0, Y: 1},
		},
		{
			testName:       "Move right from (1,1)",
			initialPos:     Coordinate2D[int]{X: 1, Y: 1},
			direction:      Right(),
			expectedNewPos: Coordinate2D[int]{X: 2, Y: 1},
		},
		{
			testName:       "Move down-left from (0,0)",
			initialPos:     Coordinate2D[int]{X: 0, Y: 0},
			direction:      DownLeft(),
			expectedNewPos: Coordinate2D[int]{X: -1, Y: -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actor := NewActor(tc.initialPos, tc.direction)
			actor.Move()
			assert.Equal(t, tc.expectedNewPos, actor.Position)
			assert.Equal(t, tc.direction, actor.Direction)
		})
	}
}

// TestActorNext tests the Next method of the Actor struct
func TestActorNext(t *testing.T) {
	testCases := []struct {
		testName        string
		initialPos      Coordinate2D[int]
		direction       Direction
		expectedNextPos Coordinate2D[int]
	}{
		{
			testName:        "Next position up from origin",
			initialPos:      Coordinate2D[int]{X: 0, Y: 0},
			direction:       Up(),
			expectedNextPos: Coordinate2D[int]{X: 0, Y: 1},
		},
		{
			testName:        "Next position right from (1,1)",
			initialPos:      Coordinate2D[int]{X: 1, Y: 1},
			direction:       Right(),
			expectedNextPos: Coordinate2D[int]{X: 2, Y: 1},
		},
		{
			testName:        "Next position down-left from (0,0)",
			initialPos:      Coordinate2D[int]{X: 0, Y: 0},
			direction:       DownLeft(),
			expectedNextPos: Coordinate2D[int]{X: -1, Y: -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			actor := NewActor(tc.initialPos, tc.direction)
			nextPos := actor.Next()
			assert.Equal(t, tc.expectedNextPos, nextPos)
			// Ensure the actor's position hasn't changed
			nextPos = Coordinate2D[int]{}
			assert.Equal(t, tc.initialPos, actor.Position)
			assert.Equal(t, tc.direction, actor.Direction)
		})
	}
}
