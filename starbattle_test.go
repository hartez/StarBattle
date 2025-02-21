package main

import (
	"testing"
)

func TestPassing(test *testing.T) {
	// This should be a passing test
}

func TestFailing(test *testing.T) {
	test.Fatalf("This should be a failing test")
}
