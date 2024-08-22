package solver

// Solver contains all the methods that a solver must implement, so that the framework can use it
// to solve a day's puzzle, and bench it.
type Solver interface {
	AddHyperParams(params ...any)
	ParsePart1(lines []string)
	ParsePart2(lines []string)
	SolvePart1() string
	SolvePart2() string
}
