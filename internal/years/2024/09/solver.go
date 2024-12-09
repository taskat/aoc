package day09

import (
	"fmt"
	"strings"

	"github.com/taskat/aoc/internal/years/2024/days"
)

// day is the day of the solver
const day = 9

// init registers the solver for day 09
func init() {
	fmt.Println("Registering day", day)
	days.AddDay(day, &Solver{})
}

// Solver implements the puzzle solver for day 09
type Solver struct{}

// AddHyperParams adds hyper parameters to the solver
func (s *Solver) AddHyperParams(params ...any) {}

// parse handle the common parsing logic for both parts
func (s *Solver) parse(lines []string) system {
	startId := 0
	blocks := make(system, len(lines[0]))
	for i, length := range lines[0] {
		len := int(length - '0')
		if i%2 == 0 {
			blocks[i] = newBlock(startId, len)
			startId++
		} else {
			blocks[i] = newEmptyBlock(len)
		}
	}
	return blocks
}

type system []block

func (s system) checksum() int {
	checksum := 0
	start := 0
	for _, b := range s {
		checksum += b.checksum(start)
		start += b.length
	}
	return checksum
}

func (s *system) compact() {
	for i := 0; i < len(*s); i++ {
		if !(*s)[i].isEmpty() {
			continue
		}
		block := (*s)[i]
		newBlocks := s.getBlocksFromEnd(block.length, i)
		s.insert(newBlocks, i)
		i += len(newBlocks) - 1
	}
}

func (s *system) compactWithoutFragmentation() {
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i].isEmpty() {
			continue
		}
		blockToMove := (*s)[i]
		space := s.getEmptyBlockIndex(blockToMove.length, i)
		if space == -1 {
			continue
		}
		emptyBlock := (*s)[space]
		newBlocks := []block{blockToMove}
		if blockToMove.length != emptyBlock.length {
			newBlocks = append(newBlocks, newEmptyBlock(emptyBlock.length-blockToMove.length))
		}
		(*s)[i] = newEmptyBlock(blockToMove.length)
		s.insert(newBlocks, space)
	}
}

func (s *system) getBlocksFromEnd(length int, dest int) []block {
	blocks := make([]block, 0, length)
	for i := len(*s) - 1; i > dest && length > 0; i-- {
		if (*s)[i].isEmpty() {
			continue
		}
		newBlock, remainings := (*s)[i].move(length)
		s.insert(remainings, i)
		blocks = append(blocks, newBlock)
		length -= newBlock.length
	}
	return blocks
}

func (s system) getEmptyBlockIndex(length, start int) int {
	for i := 0; i < start; i++ {
		if s[i].isEmpty() && s[i].length >= length {
			return i
		}
	}
	return -1
}

func (s *system) insert(blocks []block, start int) {
	newSystem := make(system, len(*s)+len(blocks)-1)
	copy(newSystem, (*s)[:start])
	copy(newSystem[start:], blocks)
	copy(newSystem[start+len(blocks):], (*s)[start+1:])
	*s = newSystem
}

func (s system) String() string {
	var sb strings.Builder
	for _, b := range s {
		if b != (block{}) {
			sb.WriteString(b.String())
		}
	}
	return sb.String()
}

type block struct {
	id     int
	length int
}

func newBlock(id, length int) block {
	return block{id: id, length: length}
}

func newEmptyBlock(length int) block {
	return block{id: -1, length: length}
}

func (b block) checksum(start int) int {
	if b.isEmpty() {
		return 0
	}
	checksum := (b.length - 1) * b.length / 2
	return (checksum + b.length*start) * b.id
}

func (b block) isEmpty() bool {
	return b.id == -1
}

func (b block) len() int {
	return b.length
}

func (b block) move(len int) (moved block, remaining []block) {
	if len >= b.length {
		return b, []block{newEmptyBlock(b.length)}
	}
	b.length -= len
	return newBlock(b.id, len), []block{b, newEmptyBlock(len)}
}

func (b block) String() string {
	if b.isEmpty() {
		return strings.Repeat(".", b.length)
	}
	return strings.Repeat(fmt.Sprintf("%d", b.id), b.length)
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
