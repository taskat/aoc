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
			blocks[i] = newFileBlock(startId, len)
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
		start += b.len()
	}
	return checksum
}

func (s *system) compact() {
	for i := 0; i < len(*s); i++ {
		if !(*s)[i].isEmpty() {
			continue
		}
		block := (*s)[i]
		newBlocks := s.getBlocksFromEnd(block.len(), i)
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
		space := s.getEmptyBlockIndex(blockToMove.len(), i)
		if space == -1 {
			continue
		}
		emptyBlock := (*s)[space]
		newBlocks := []block{blockToMove}
		if blockToMove.len() != emptyBlock.len() {
			newBlocks = append(newBlocks, newEmptyBlock(emptyBlock.len()-blockToMove.len()))
		}
		(*s)[i] = newEmptyBlock(blockToMove.len())
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
		length -= newBlock.len()
	}
	return blocks
}

func (s system) getEmptyBlockIndex(length, start int) int {
	for i := 0; i < start; i++ {
		if s[i].isEmpty() && s[i].len() >= length {
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
		if b != nil {
			sb.WriteString(b.String())
		}
	}
	return sb.String()
}

type block interface {
	checksum(start int) int
	getId() int
	isEmpty() bool
	len() int
	move(len int) (block, []block)
	fmt.Stringer
}

type fileBlock struct {
	id     int
	length int
}

func newFileBlock(id, length int) *fileBlock {
	return &fileBlock{id: id, length: length}
}

func (b *fileBlock) checksum(start int) int {
	checksum := (b.length - 1) * b.length / 2
	return (checksum + b.length*start) * b.id
}

func (b *fileBlock) getId() int {
	return b.id
}

func (b *fileBlock) isEmpty() bool {
	return false
}

func (b *fileBlock) len() int {
	return b.length
}

func (b *fileBlock) move(len int) (moved block, remaining []block) {
	if len >= b.length {
		return b, []block{newEmptyBlock(b.length)}
	}
	b.length -= len
	return newFileBlock(b.id, len), []block{b, newEmptyBlock(len)}
}

func (b *fileBlock) String() string {
	return strings.Repeat(fmt.Sprintf("%d", b.id), b.length)
}

type emptyBlock struct {
	length int
}

func newEmptyBlock(length int) emptyBlock {
	return emptyBlock{length: length}
}

func (b emptyBlock) checksum(start int) int {
	return 0
}

func (b emptyBlock) getId() int {
	return 0
}

func (b emptyBlock) isEmpty() bool {
	return true
}

func (b emptyBlock) len() int {
	return b.length
}

func (b emptyBlock) move(len int) (block, []block) {
	panic("Empty block cannot be moved")
}

func (b emptyBlock) String() string {
	return strings.Repeat(".", b.length)
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
