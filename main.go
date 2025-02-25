package main

import (
	"fmt"
	"os"
)

func main() {
	puzzleFile := "5_1.txt"

	if len(os.Args) > 1 {
		puzzleFile = os.Args[1]
	}

	fmt.Printf("Working on %s\n", puzzleFile)

	board := Parse(puzzleFile)

	fmt.Print(board)
}
