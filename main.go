package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"
	"sync"
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

	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()

	fmt.Printf("CPUs: %d; Max processes: %d\n", numCPU, maxProcs)
	fmt.Printf("Working on %s\n", puzzleFile)
	fmt.Print(board)

	solutionChannel := make(chan *Board)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)

	start := time.Now()

	go board.SolveParallel(solutionChannel, &wg, ctx)

	go func() {
		wg.Wait()
		solutionChannel <- nil
	}()

	solution := <-solutionChannel
	end := time.Now()

	if solution == nil {
		fmt.Printf("Could not find a solution\n")
	} else {
		cancel()
		fmt.Printf("Solved!\n")
		fmt.Print(solution)
	}

	fmt.Printf("Solve time was %s\n", end.Sub(start))
}
