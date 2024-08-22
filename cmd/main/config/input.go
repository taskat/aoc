package config

import (
	"fmt"
)

// InputType is a type that represents the type of input that the solver will receive
type InputType interface {
	// isInputType is an empty method that makes the interface private (unimplementable from other packages)
	isInputType()
	fmt.Stringer
}

// parseInputType parses the input type from the command line arguments
// It returns nil if the input type is not valid, otherwise it returns the input type
// which can be either RealInput or TestInput
func parseInputType(inputType string) InputType {
	switch inputType {
	case "real":
		return RealInput{}
	default:
		return parseTestInput(inputType)
	}
}

// RealInput is the input type that represents the real input
type RealInput struct{}

// isInputType is a method that is used to make RealInput implement the InputType interface
func (RealInput) isInputType() {}

// String is a method that returns a string representation of RealInput
func (RealInput) String() string {
	return "real"
}

// TestInput is the input type that represents a test input
// It is a number and is used to specify which test input to use
type TestInput int

// parseTestInput parses the test input from the command line arguments
// It returns nil if the input type is not a test input
func parseTestInput(inputType string) InputType {
	var testInput TestInput
	_, err := fmt.Sscanf(inputType, "test-%d", &testInput)
	if err != nil {
		return nil
	}
	return testInput
}

// isInputType is a method that is used to make TestInput implement the InputType interface
func (TestInput) isInputType() {}

// String is a method that returns a string representation of TestInput
func (t TestInput) String() string {
	return fmt.Sprintf("test-%d", t)
}
