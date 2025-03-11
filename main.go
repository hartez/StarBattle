package main

import (
	"fmt"
	"log"
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

	_, solution, err := board.Solve()

	end := time.Now()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solved!\n")
	fmt.Print(solution)
	fmt.Printf("Solve time was %s\n", end.Sub(start))
}
