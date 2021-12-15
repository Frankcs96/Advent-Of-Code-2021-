package day15_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day15"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("example_input.txt")
	expected := 40

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("example_input.txt")
	expected := 315

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 441

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input.txt")
	expected := 2849

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
