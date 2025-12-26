package coordinate

import (
	"github.com/taskat/aoc/pkg/utils/containers/set"
	"github.com/taskat/aoc/pkg/utils/slices"
	"github.com/taskat/aoc/pkg/utils/types"
)

// Compression2D represents a 2D compression mapping for 2D coordinates
type Compression2D[T types.Real] struct {
	xMapping map[T]int
	yMapping map[T]int
	inverseX map[int]T
	inverseY map[int]T
}

// NewCompression2D creates a new Compression2D instance
func NewCompression2D[T types.Real]() *Compression2D[T] {
	return &Compression2D[T]{
		xMapping: make(map[T]int),
		yMapping: make(map[T]int),
		inverseX: make(map[int]T),
		inverseY: make(map[int]T),
	}
}

// Compress compresses a coordinate
func (c *Compression2D[T]) Compress(coord Coordinate2D[T]) Coordinate2D[int] {
	return NewCoordinate2D(
		c.xMapping[coord.X],
		c.yMapping[coord.Y],
	)
}

// Decompress decompresses a coordinate
func (c *Compression2D[T]) Decompress(coord Coordinate2D[int]) Coordinate2D[T] {
	return NewCoordinate2D(
		c.inverseX[coord.X],
		c.inverseY[coord.Y],
	)
}

// compress compresses the given array of values and returns the mapping and the inverse mapping
func compress[T types.Real](values []T) (map[T]int, map[int]T) {
	valueSet := set.FromSlice(values)
	uniqueValues := valueSet.ToSlice()
	slices.Sort(uniqueValues, func(a, b T) bool { return a < b })
	mapping := make(map[T]int)
	inverseMapping := make(map[int]T)
	for i, v := range uniqueValues {
		mapping[v] = i * 2
		inverseMapping[i*2] = v
	}
	return mapping, inverseMapping
}

// Compress creates a compression mapping from the given coordinates
// and compresses the coordinates
func Compress2D[T types.Real](coords []Coordinate2D[T]) ([]Coordinate2D[int], *Compression2D[T]) {
	xMapping, inverseX := compress(slices.Map(coords, func(c Coordinate2D[T]) T { return c.X }))
	yMapping, inverseY := compress(slices.Map(coords, func(c Coordinate2D[T]) T { return c.Y }))
	compression := &Compression2D[T]{
		xMapping: xMapping,
		yMapping: yMapping,
		inverseX: inverseX,
		inverseY: inverseY,
	}
	compressedCoords := make([]Coordinate2D[int], len(coords))
	for i, c := range coords {
		compressedCoords[i] = NewCoordinate2D(compression.xMapping[c.X], compression.yMapping[c.Y])
	}

	return compressedCoords, compression
}
