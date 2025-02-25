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

func (board Board) index(row int, col int) (int, error) {

	boardSize := board.size

	if row < 0 || col < 0 || row >= boardSize || col >= boardSize {
		// Out of bounds
		return -1, fmt.Errorf("%d, %d is out of bounds", row, col)
	}

	return boardSize*row + col, nil
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
