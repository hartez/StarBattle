package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	parallel := flag.Bool("p", false, "solve puzzle in parallel")
	puzzleFile := flag.String("i", "5_1.txt", "puzzle input file")

	flag.Parse()

	if *parallel {
		SolveInParallel(*puzzleFile)
	} else {
		SolveSequential(*puzzleFile)
	}
}

func SolveSequential(puzzleFile string) {
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

func SolveInParallel(puzzleFile string) {
	board := Parse(puzzleFile)

	solutionChannel := make(chan Board)

	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()

	fmt.Printf("CPUs: %d; Max processes: %d\n", numCPU, maxProcs)
	fmt.Printf("Working on %s\n", puzzleFile)
	fmt.Print(board)

	start := time.Now()

	go board.SolveParallel(solutionChannel)

	solution := <-solutionChannel

	end := time.Now()

	fmt.Printf("Solved!\n")
	fmt.Print(solution)

	fmt.Printf("Solve time was %s\n", end.Sub(start))
}
