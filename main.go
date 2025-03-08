package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	puzzleFile := "5_1.txt"

	if len(os.Args) > 1 {
		puzzleFile = os.Args[1]
	}

	board := Parse(puzzleFile)

	fmt.Printf("Working on %s\n", puzzleFile)
	fmt.Print(board)

	_, solution, err := board.Solve()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solved!\n")
	fmt.Print(solution)
}
