package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Board struct {
	size    int
	stars   int
	squares []Square
	regions [][]int
}

func newBoard(size int, stars int) Board {
	var board Board

	board.size = size
	board.stars = stars
	board.squares = make([]Square, size*size)
	board.regions = make([][]int, size)

	return board
}

func Parse(fileName string) Board {

	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	data := string(content)

	// Check for a stars specification at the beginning
	stars := 1
	before, after, found := strings.Cut(data, "*")

	if found {
		// There's a star spec at the beginning
		stars, err = strconv.Atoi(before)
		if err != nil {
			// Handle this, the star spec was bad
			log.Fatalf("Bad star specification: %s", err)
		}

		data = after
	}

	squares := strings.Split(data, ",")

	length := (float64)(len(squares))

	size := int(math.Sqrt(length))

	board := newBoard(size, stars)

	for index, region := range squares {
		if region, err := strconv.Atoi(strings.TrimSpace(region)); err == nil {
			board.setRegion(index, region-1)
		} else {
			fmt.Println(err)
		}
	}

	return board
}

func (board Board) String() string {

	var sb strings.Builder

	size := board.size

	plural := "s"
	if board.stars == 1 {
		plural = ""
	}

	sb.WriteString(fmt.Sprintf("%d x %d, %d star%s\n", size, size, board.stars, plural))

	for row := range size {
		for col := range size {

			value, err := board.value(row, col)

			if err != nil {
				return err.Error()
			}

			region, err := board.region(row, col)

			if err != nil {
				return err.Error()
			}

			sb.WriteString(sectionColor(region))
			sb.WriteString("[")
			sb.WriteString(value.String())
			sb.WriteString("]")
			sb.WriteString("\033[0m")
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func sectionColor(section int) string {

	modification := 0

	if section > 7 {
		modification = 1 // bold
	} else if section > 14 {
		modification = 4 // underline
	}

	return fmt.Sprintf("\033[%d;%dm", modification, ((section % 7) + 31))
}

func (board Board) findEmptySquare() (row int, column int, err error) {

	boardSize := board.size

	for index, square := range board.squares {
		if square == UNKNOWN {
			return index / boardSize, index % boardSize, nil
		}
	}

	return -1, -1, fmt.Errorf("no more available squares")
}

func (original Board) copy() Board {
	board := newBoard(original.size, original.stars)
	copy(board.squares, original.squares)
	copy(board.regions, original.regions)
	return board
}

func (board Board) Solve() (bool, Board, error) {

	if !board.isValid() {
		return false, board, nil
	}

	if board.hasEnoughStars() {
		return true, board, nil
	}

	nextRow, nextCol, err := board.findEmptySquare()

	if err != nil {
		return false, board, fmt.Errorf("no solution found")
	}

	nextBoard := board.copy()

	nextBoard.setValue(nextRow, nextCol, STAR)

	solved, solvedBoard, err := nextBoard.Solve()

	if err != nil {
		return false, board, fmt.Errorf("no solution found")
	}

	if solved {
		return true, solvedBoard, nil
	}

	nextBoard = board.copy()
	nextBoard.setValue(nextRow, nextCol, NOTSTAR)

	return nextBoard.Solve()
}

func (board Board) index(row int, col int) (int, error) {

	boardSize := board.size

	if row < 0 || col < 0 || row >= boardSize || col >= boardSize {
		// Out of bounds
		return -1, fmt.Errorf("%d, %d is out of bounds", row, col)
	}

	return boardSize*row + col, nil
}

func (board Board) value(row int, column int) (Square, error) {
	index, err := board.index(row, column)

	if err != nil {
		return -1, err
	}

	return board.squares[index], nil
}

func (board Board) setValue(row int, column int, value Square) error {
	index, err := board.index(row, column)

	if err != nil {
		return err
	}

	board.squares[index] = value
	return nil
}

func (board Board) region(row int, col int) (int, error) {

	index, err := board.index(row, col)

	if err != nil {
		return -1, err
	}

	for regionIndex, regionMap := range board.regions {
		if slices.Contains(regionMap, index) {
			return regionIndex, nil
		}
	}

	return -1, fmt.Errorf("region not found for %d, %d", row, col)
}

func (board Board) setRegion(index int, region int) {
	if board.regions[region] == nil {
		board.regions[region] = make([]int, 1, 10)
		board.regions[region][0] = index
	} else {
		board.regions[region] = append(board.regions[region], index)
	}
}

func (board Board) hasEnoughStars() bool {
	var starCount int

	for _, square := range board.squares {
		if square == STAR {
			starCount += 1
		}
	}

	return starCount == (board.size * board.stars)
}

func (board Board) isValid() bool {
	boardSize := board.size
	requiredStars := board.stars

	for row := range boardSize {
		stars, notStars := board.countInRow(row)

		if stars > requiredStars {
			// Too many stars is invalid
			return false
		}

		if stars+notStars == boardSize && stars < requiredStars {
			// If the row is full and we don't have enough stars, the board is invalid
			return false
		}
	}

	for column := range boardSize {
		stars, notStars := board.countInColumn(column)

		if stars > requiredStars {
			return false
		}

		if stars+notStars == boardSize && stars < requiredStars {
			return false
		}
	}

	for section := range boardSize {

		stars, notStars, regionSize := board.countInRegion(section)

		if stars > requiredStars {
			return false
		}

		if stars+notStars == regionSize && stars < requiredStars {
			return false
		}
	}

	for index, square := range board.squares {
		if square == STAR {
			if board.anyAdjacentStars(index/boardSize, index%boardSize) {
				return false
			}
		}
	}

	return true
}

func (board Board) countInRow(row int) (stars int, notStars int) {
	start := row * board.size
	end := start + board.size

	for _, value := range board.squares[start:end] {
		if value == STAR {
			stars += 1
		}

		if value == NOTSTAR {
			notStars += 1
		}
	}

	return
}

func (board Board) countInColumn(column int) (stars int, notStars int) {
	boardSize := board.size
	squares := board.squares
	end := (boardSize * boardSize) - (boardSize - column)

	for squareIndex := column; squareIndex <= end; squareIndex += boardSize {
		value := squares[squareIndex]

		if value == STAR {
			stars += 1
		}

		if value == NOTSTAR {
			notStars += 1
		}
	}

	return
}

func (board Board) countInRegion(region int) (stars int, notStars int, size int) {
	regionMap := board.regions[region]
	size = len(regionMap)
	squares := board.squares

	for _, squareIndex := range regionMap {

		value := squares[squareIndex]

		if value == STAR {
			stars += 1
		}

		if value == NOTSTAR {
			notStars += 1
		}
	}

	return
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
