package main

import (
	"testing"
)

func TestBenchmarkSmall(test *testing.T) {
	board := Parse("5_1.txt")
	err := board.solveAndVerify()

	if err != nil {
		test.Fatal(err)
	}
}

func TestBenchmarkMedium(test *testing.T) {
	board := Parse("10_2.txt")
	err := board.solveAndVerify()

	if err != nil {
		test.Fatal(err)
	}
}

func TestBenchmarkHard(test *testing.T) {
	board := Parse("14_3_2.txt")
	err := board.solveAndVerify()

	if err != nil {
		test.Fatal(err)
	}
}
