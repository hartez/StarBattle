package main

import (
	"testing"
)

func TestCanParseSimpleBoardFromFile(test *testing.T) {
	testFile := "5_1.txt"

	expectedSize := 5
	expectedStars := 1

	board := Parse(testFile)

	if board.size != expectedSize {
		test.Fatalf("Board size should be %d in %s, but found %d", expectedSize, testFile, board.size)
	}

	if board.stars != expectedStars {
		test.Fatalf("Board stars should be %d in %s, but found %d", expectedStars, testFile, board.stars)
	}

	if len(board.regions) != expectedSize {
		test.Fatalf("Board should have %d regions, but found %d", expectedSize, len(board.regions))
	}
}

func TestCanParseComplexBoardFromFile(test *testing.T) {
	testFile := "14_3.txt"

	expectedSize := 14
	expectedStars := 3

	board := Parse(testFile)

	if board.size != expectedSize {
		test.Fatalf("Board size should be %d in %s, but found %d", expectedSize, testFile, board.size)
	}

	if board.stars != expectedStars {
		test.Fatalf("Board stars should be %d in %s, but found %d", expectedStars, testFile, board.stars)
	}

	if len(board.regions) != expectedSize {
		test.Fatalf("Board should have %d regions, but found %d", expectedSize, len(board.regions))
	}
}

func TestRegionMappingsFromFile(test *testing.T) {
	testFile := "5_1.txt"

	board := Parse(testFile)

	// Spot check some region mappings
	row := 0
	col := 0
	expectedRegion := 0
	actualRegion, _ := board.region(row, col)

	if actualRegion != expectedRegion {
		test.Fatalf("Expected square at %d, %d to be in region %d, but found %d", row, col, expectedRegion, actualRegion)
	}

	row = 4
	col = 4
	expectedRegion = 3
	actualRegion, _ = board.region(row, col)

	if actualRegion != expectedRegion {
		test.Fatalf("Expected square at %d, %d to be in region %d, but found %d", row, col, expectedRegion, actualRegion)
	}
}
