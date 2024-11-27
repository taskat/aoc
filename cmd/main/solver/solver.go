package solver

// Solver contains all the methods that a solver must implement, so that the framework can use it
// to solve a day's puzzle, and bench it.
type Solver interface {
	AddHyperParams(params ...any)
	SolvePart1(lines []string) string
	SolvePart2(lines []string) string
}
