package main

import (
	"fmt"
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

	regions := make([]bool, board.size)

	for row := 0; row < board.size; row++ {
		for column := 0; column < board.size; column++ {
			region, _ := board.region(row, column)
			regions[region] = true
		}
	}

	for i := 0; i < len(regions); i++ {
		if !regions[i] {
			test.Fatalf("Missing expected region %d", i)
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

func TestHasEnoughStars(test *testing.T) {
	testFile := "5_1.txt"
	board := Parse(testFile)

	if board.hasEnoughStars() {
		test.Fatalf("Board is empty")
	}

	board.setValue(0, 0, STAR)
	board.setValue(0, 1, STAR)
	board.setValue(0, 2, STAR)
	board.setValue(0, 3, STAR)
	board.setValue(0, 4, STAR)

	if !board.hasEnoughStars() {
		test.Fatalf("Board has correct number of stars")
	}
}

func TestIsInvalidIfTooManyStarsInRow(test *testing.T) {
	board := Parse("5_1.txt")

	board.setValue(2, 3, STAR)
	board.setValue(2, 4, STAR)

	if board.isValid() {
		test.Fatalf("Board should be invalid because there are too many stars in the same row")
	}
}

func TestIsInvalidIfTooManyStarsInColumn(test *testing.T) {
	board := Parse("5_1.txt")

	board.setValue(1, 2, STAR)
	board.setValue(3, 2, STAR)

	if board.isValid() {
		test.Fatalf("Board should be invalid because there are too many stars in the same column")
	}
}

func TestIsInvalidIfTooManyStarsInRegion(test *testing.T) {
	board := Parse("5_1.txt")

	board.setValue(0, 0, STAR)
	board.setValue(2, 1, STAR)

	if board.isValid() {
		test.Fatalf("Board should be invalid because there are too many stars in the same region")
	}
}

func TestIsInvalidIfRegionFullAndStarsLessThanLimit(test *testing.T) {
	board := Parse("10_2.txt")

	board.setValue(0, 8, NOTSTAR)
	board.setValue(0, 9, NOTSTAR)

	board.setValue(1, 7, NOTSTAR)
	board.setValue(1, 8, NOTSTAR)
	board.setValue(1, 9, NOTSTAR)

	board.setValue(2, 7, NOTSTAR)
	board.setValue(2, 8, STAR)

	if board.isValid() {
		fmt.Println(board)
		test.Fatalf("Board should be invalid, region 3 is full and only has one star.")
	}
}

func TestIsInvalidIfRowFullAndStarsLessThanLimit(test *testing.T) {
	board := Parse("10_2.txt")

	board.setValue(0, 0, STAR)
	board.setValue(0, 1, NOTSTAR)
	board.setValue(0, 2, NOTSTAR)
	board.setValue(0, 3, NOTSTAR)
	board.setValue(0, 4, NOTSTAR)
	board.setValue(0, 5, NOTSTAR)
	board.setValue(0, 6, NOTSTAR)
	board.setValue(0, 7, NOTSTAR)
	board.setValue(0, 8, NOTSTAR)
	board.setValue(0, 9, NOTSTAR)

	if board.isValid() {
		fmt.Println(board)
		test.Fatalf("Board should be invalid, row 0 is full and only has one star.")
	}
}

func TestIsInvalidIfColumnFullAndStarsLessThanLimit(test *testing.T) {
	board := Parse("10_2.txt")

	board.setValue(0, 0, STAR)
	board.setValue(1, 0, NOTSTAR)
	board.setValue(2, 0, NOTSTAR)
	board.setValue(3, 0, NOTSTAR)
	board.setValue(4, 0, NOTSTAR)
	board.setValue(5, 0, NOTSTAR)
	board.setValue(6, 0, NOTSTAR)
	board.setValue(7, 0, NOTSTAR)
	board.setValue(8, 0, NOTSTAR)
	board.setValue(9, 0, NOTSTAR)

	if board.isValid() {
		fmt.Println(board)
		test.Fatalf("Board should be invalid, column 0 is full and only has one star.")
	}
}

func TestIsInvalidIfEntireRowIsMarkedNotStar(test *testing.T) {
	board := Parse("5_1.txt")

	board.setValue(2, 0, NOTSTAR)
	board.setValue(2, 1, NOTSTAR)
	board.setValue(2, 2, NOTSTAR)
	board.setValue(2, 3, NOTSTAR)
	board.setValue(2, 4, NOTSTAR)

	if board.isValid() {
		test.Fatalf("Board should be invalid because the row is filled and has no star")
	}
}

func TestIsInvalidIfEntireColumnIsMarkedNotStar(test *testing.T) {
	board := Parse("5_1.txt")

	board.setValue(0, 2, NOTSTAR)
	board.setValue(1, 2, NOTSTAR)
	board.setValue(2, 2, NOTSTAR)
	board.setValue(3, 2, NOTSTAR)
	board.setValue(4, 2, NOTSTAR)

	if board.isValid() {
		test.Fatalf("Board should be invalid because the column is filled and has no star")
	}
}

func TestIsInvalidIfEntireRegionIsMarkedNotStar(test *testing.T) {
	board := Parse("5_1.txt")

	board.setValue(4, 2, NOTSTAR)
	board.setValue(4, 3, NOTSTAR)

	if board.isValid() {
		test.Fatalf("Board should be invalid because the region is filled and has no star")
	}
}

func TestIsValidIfStarsInRowLessThanPuzzleLimit(test *testing.T) {
	board := Parse("10_2_2.txt")

	board.setValue(1, 0, STAR)
	board.setValue(1, 2, STAR)

	if !board.isValid() {
		test.Fatalf("Board should be valid, 2 non-adjacent stars are allowed in the same row in this puzzle.")
	}
}

func TestIsValidIfStarsInColumnLessThanPuzzleLimit(test *testing.T) {
	board := Parse("10_2_2.txt")

	board.setValue(0, 0, STAR)
	board.setValue(3, 0, STAR)

	if !board.isValid() {
		test.Fatalf("Board should be valid, 2 non-adjacent stars are allowed in the same column in this puzzle.")
	}
}

func TestIsValidIfStarsInRegionLessThanPuzzleLimit(test *testing.T) {
	board := Parse("10_2.txt")

	board.setValue(0, 0, STAR)
	board.setValue(3, 0, STAR)

	if !board.isValid() {
		test.Fatalf("Board should be valid, 2 non-adjacent stars are allowed in the same region in this puzzle.")
	}
}

func (board Board) solveAndVerify() error {
	isSolved, solvedBoard, err := board.Solve()

	if !isSolved || err != nil {
		return fmt.Errorf("board was not solved")
	}

	if !solvedBoard.ensureNoAdjacentStars() {
		return fmt.Errorf("solution was invalid, adjacent stars found")
	}

	return nil
}

func (board Board) ensureNoAdjacentStars() bool {

	boardSize := board.size

	for index, square := range board.squares {
		if square == STAR {
			if board.anyAdjacentStars(index/boardSize, index%boardSize) {
				return false
			}
		}
	}

	return true
}

func (board Board) anyAdjacentStars(row int, column int) bool {
	squares := board.squares

	for r := -1; r <= 1; r++ {
		for c := -1; c < 1; c++ {

			if r == 0 && c == 0 {
				continue
			}

			index, err := board.index(row+r, column+c)

			if err != nil {
				// The row/column is off the edge of the board, ignore it
				continue
			}

			if squares[index] == STAR {
				return true
			}
		}
	}

	return false
}
