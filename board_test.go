package main

import (
	"testing"
)

func TestCanParseSimpleBoardFromFile(test *testing.T) {
	testFile := "5_1.txt"

	expectedSize := 5
	expectedStars := 1

	checkExpectedValues(expectedSize, expectedStars, testFile, test)
}

func TestCanParseComplexBoardFromFile(test *testing.T) {
	testFile := "14_3.txt"

	expectedSize := 14
	expectedStars := 3

	checkExpectedValues(expectedSize, expectedStars, testFile, test)
}

func TestRegionMappingsFromFile(test *testing.T) {
	testFile := "5_1.txt"

	board := Parse(testFile)

	// Spot check some region mappings
	checkRegion(board, 0, 0, 0, test)
	checkRegion(board, 4, 4, 3, test)
}

func checkRegion(board Board, row int, col int, expectedRegion int, test *testing.T) {
	actualRegion, _ := board.region(row, col)

	if actualRegion != expectedRegion {
		test.Fatalf("Expected square at %d, %d to be in region %d, but found %d", row, col, expectedRegion, actualRegion)
	}
}

func checkExpectedRegions(board Board, test *testing.T) {

	sections := make([]bool, board.size)

	for row := 0; row < board.size; row++ {
		for column := 0; column < board.size; column++ {
			region, _ := board.region(row, column)
			sections[region] = true
		}
	}

	for i := 0; i < len(sections); i++ {
		if !sections[i] {
			test.Fatalf("Missing expected section %d", i)
		}
	}
}

func checkExpectedValues(expectedSize int, expectedStars int, testFile string, test *testing.T) {
	board := Parse(testFile)

	if board.size != expectedSize {
		test.Fatalf("Board size should be %d in %s, but found %d", expectedSize, testFile, board.size)
	}

	if board.stars != expectedStars {
		test.Fatalf("Board stars should be %d in %s, but found %d", expectedStars, testFile, board.stars)
	}

	checkExpectedRegions(board, test)
}
