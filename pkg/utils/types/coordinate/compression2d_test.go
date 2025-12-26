package coordinate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompression2D(t *testing.T) {
	testCases := []struct {
		testName            string
		coords              []Coordinate2D[int]
		expectedCompressed  []Coordinate2D[int]
		expectedCompression *Compression2D[int]
	}{
		{
			testName: "Simple case",
			coords: []Coordinate2D[int]{
				NewCoordinate2D(10, 20),
				NewCoordinate2D(30, 40),
				NewCoordinate2D(10, 40),
			},
			expectedCompressed: []Coordinate2D[int]{
				NewCoordinate2D(0, 0),
				NewCoordinate2D(2, 2),
				NewCoordinate2D(0, 2),
			},
			expectedCompression: &Compression2D[int]{
				xMapping: map[int]int{10: 0, 30: 2},
				yMapping: map[int]int{20: 0, 40: 2},
				inverseX: map[int]int{0: 10, 2: 30},
				inverseY: map[int]int{0: 20, 2: 40},
			},
		},
		{
			testName: "With negative coordinates",
			coords: []Coordinate2D[int]{
				NewCoordinate2D(-5, -10),
				NewCoordinate2D(0, 0),
				NewCoordinate2D(-5, 0),
			},
			expectedCompressed: []Coordinate2D[int]{
				NewCoordinate2D(0, 0),
				NewCoordinate2D(2, 2),
				NewCoordinate2D(0, 2),
			},
			expectedCompression: &Compression2D[int]{
				xMapping: map[int]int{-5: 0, 0: 2},
				yMapping: map[int]int{-10: 0, 0: 2},
				inverseX: map[int]int{0: -5, 2: 0},
				inverseY: map[int]int{0: -10, 2: 0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			compressed, compression := Compress2D(tc.coords)
			assert.Equal(t, tc.expectedCompressed, compressed)
			assert.Equal(t, tc.expectedCompression.xMapping, compression.xMapping)
			assert.Equal(t, tc.expectedCompression.yMapping, compression.yMapping)
			assert.Equal(t, tc.expectedCompression.inverseX, compression.inverseX)
			assert.Equal(t, tc.expectedCompression.inverseY, compression.inverseY)
		})
	}
}

func TestCompress(t *testing.T) {
	testCases := []struct {
		testName           string
		compression        *Compression2D[int]
		c                  Coordinate2D[int]
		expectedCompressed Coordinate2D[int]
	}{
		{
			testName: "Simple compression",
			compression: &Compression2D[int]{
				xMapping: map[int]int{10: 0, 20: 2},
				yMapping: map[int]int{30: 0, 40: 2},
			},
			c:                  NewCoordinate2D(10, 40),
			expectedCompressed: NewCoordinate2D(0, 2),
		},
		{
			testName: "Another compression",
			compression: &Compression2D[int]{
				xMapping: map[int]int{-5: 0, 0: 2},
				yMapping: map[int]int{-10: 0, 0: 2},
			},
			c:                  NewCoordinate2D(0, -10),
			expectedCompressed: NewCoordinate2D(2, 0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			compressed := tc.compression.Compress(tc.c)
			assert.Equal(t, tc.expectedCompressed, compressed)
		})
	}
}

func TestDecompress(t *testing.T) {
	testCases := []struct {
		testName             string
		compression          *Compression2D[int]
		c                    Coordinate2D[int]
		expectedDecompressed Coordinate2D[int]
	}{
		{
			testName: "Simple decompression",
			compression: &Compression2D[int]{
				inverseX: map[int]int{0: 10, 2: 20},
				inverseY: map[int]int{0: 30, 2: 40},
			},
			c:                    NewCoordinate2D(2, 0),
			expectedDecompressed: NewCoordinate2D(20, 30),
		},
		{
			testName: "Another decompression",
			compression: &Compression2D[int]{
				inverseX: map[int]int{0: -5, 2: 0},
				inverseY: map[int]int{0: -10, 2: 0},
			},
			c:                    NewCoordinate2D(0, 2),
			expectedDecompressed: NewCoordinate2D(-5, 0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			decompressed := tc.compression.Decompress(tc.c)
			assert.Equal(t, tc.expectedDecompressed, decompressed)
		})
	}
}
