package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	puzzleFile := "5_1.txt"

	if len(os.Args) > 1 {
		puzzleFile = os.Args[1]
	}

	board := Parse(puzzleFile)

	fmt.Printf("Working on %s\n", puzzleFile)
	fmt.Print(board)

	start := time.Now()

	isSolved, solution := board.Solve()

	end := time.Now()

	if isSolved {
		fmt.Printf("Solved!\n")
		fmt.Print(solution)
	} else {
		fmt.Printf("Puzzle could not be solved.\n")
	}

	fmt.Printf("Solve time was %s\n", end.Sub(start))
}
