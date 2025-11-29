package day09

import (
	"fmt"

	"github.com/taskat/aoc/internal/years/2024/days"
)

// day is the day of the solver
const day = 9

// init registers the solver for day 09
func init() {
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 09
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...string) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) system {
	currentId := 0
	system := make(system, 0, len(lines[0])*5)
	for i, length := range lines[0] {
		len := int(length - '0')
		if i%2 == 0 {
			for range len {
				system = append(system, currentId)
			}
			currentId++
		} else {
			for range len {
				system = append(system, -1)
			}
		}
	}
	return system
}

// system represents the memory blocks system
type system []int

// checksum computes the checksum of the system
// It is calculated by summing the checksums of each block
func (s system) checksum() int {
	checksum := 0
	for i, id := range s {
		if id == -1 {
			continue
		}
		checksum += id * i
	}
	return checksum
}

// compact compacts the system by moving blocks to fill empty spaces
// if necessary it breaks block from the back to fill smaller empty spaces
func (s system) compact() {
	start := 0
	for {
		emptyBlock, ok := s.getEmptyBlockBetween(start, len(s))
		if !ok {
			break
		}
		blockToMove, ok := s.getBlockFromBack(len(s) - 1)
		if !ok {
			break
		}
		if blockToMove.start < emptyBlock.start {
			break
		}
		if blockToMove.length <= emptyBlock.length {
			s.move(blockToMove, emptyBlock)
			start = emptyBlock.start + blockToMove.length
			continue
		}
		movableStart := blockToMove.start + (blockToMove.length - emptyBlock.length)
		movableBlock := newBlock(movableStart, emptyBlock.length)
		s.move(movableBlock, emptyBlock)
		start = emptyBlock.start + emptyBlock.length
	}
}

// compactWithoutFragmentation compacts the system without breaking blocks
func (s system) compactWithoutFragmentation() {
	end := len(s) - 1
	for {
		blockToMove, ok := s.getBlockFromBack(end)
		if !ok {
			break
		}
		end = blockToMove.start - 1
		emptyBlock, ok := s.getEmptyBlockBetweenWithLength(0, blockToMove.start, blockToMove.length)
		if !ok {
			continue
		}
		s.move(blockToMove, emptyBlock)
	}
}

// getBlockFromBack finds the first block from the back that is not empty
func (s system) getBlockFromBack(start int) (block, bool) {
	length := 0
	blockStart := 0
	blockId := -1
	for i := start; i >= 0; i-- {
		if s[i] == -1 {
			if length > 0 {
				return newBlock(blockStart, length), true
			}
			continue
		}
		if s[i] != blockId && blockId != -1 {
			return newBlock(blockStart, length), true
		}
		if blockId == -1 {
			blockId = s[i]
		}
		length++
		blockStart = i
	}
	if length > 0 {
		return newBlock(blockStart, length), true
	}
	return block{}, false
}

// getEmptyBlockBetween finds the first empty block from the given index
func (s system) getEmptyBlockBetween(start, end int) (block, bool) {
	length := 0
	blockStart := 0
	for i := start; i < end; i++ {
		if s[i] != -1 {
			if length > 0 {
				return newBlock(blockStart, length), true
			}
			continue
		}
		if blockStart == 0 {
			blockStart = i
		}
		length++
	}
	if length > 0 {
		return newBlock(blockStart, length), true
	}
	return block{}, false
}

// getEmptyBlockBetween finds the first empty block from the given index
func (s system) getEmptyBlockBetweenWithLength(start, end, expectedLength int) (block, bool) {
	for i := start; i < end; i++ {
		if s[i] != -1 {
			continue
		}
		valid := true
		for offSet := 1; offSet < expectedLength; offSet++ {
			if i+offSet >= len(s) || s[i+offSet] != -1 {
				valid = false
				i += offSet
				break
			}
		}
		if valid {
			return newBlock(i, expectedLength), true
		}
	}
	return block{}, false
}

// move moves a block from one position to another
// If the block to move is larger than the destination block,
// it panics. It will automatically fill the source block with empty values.
func (s system) move(from, to block) {
	if from.length > to.length {
		panic("cannot move larger block to smaller block")
	}
	for i := to.start; i < to.start+from.length; i++ {
		offset := i - to.start
		s[i] = s[from.start+offset]
		s[from.start+offset] = -1
	}
}

// String returns the string representation of the system
func (s system) String() string {
	str := ""
	for _, id := range s {
		if id == -1 {
			str += "."
			continue
		}
		str += fmt.Sprintf("%d", id)
	}
	return str
}

// block represents a block in the system
type block struct {
	start  int
	length int
}

// newBlock creates a new block
func newBlock(start, length int) block {
	return block{start: start, length: length}
}

// SolvePart1 solves part 1 of the puzzle
func (s *Solver) SolvePart1(lines []string) string {
	system := s.parse(lines)
	system.compact()
	checksum := system.checksum()
	return fmt.Sprintf("%d", checksum)
}

// SolvePart2 solves part 2 of the puzzle
func (s *Solver) SolvePart2(lines []string) string {
	system := s.parse(lines)
	system.compactWithoutFragmentation()
	checksum := system.checksum()
	return fmt.Sprintf("%d", checksum)
}
