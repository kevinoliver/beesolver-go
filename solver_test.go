package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	dictionary := NewDictionaryFromReader(strings.NewReader("ogselmd\ndoom\ndogs\ndotty\nseem\n"))
	puzzle, err := NewPuzzle('d', "ogselm")
	if err != nil {
		t.Fatal()
	}
	solver := Solver{dictionary, puzzle}
	results := solver.Solve()
	if len(results) != 3 {
		t.Fatalf("Expected 3 results, got %d", len(results))
	}
	if results[0].Word() != "dogs" {
		t.Fatalf("Expected 'dogs' got '%s'", results[0].Word())
	}
	if results[1].Word() != "doom" {
		t.Fatalf("Expected 'doom' got '%s'", results[1].Word())
	}
	res2 := results[2]
	if res2.Word() != "ogselmd" {
		t.Fatalf("Expected 'ogselmd' got '%s'", res2.Word())
	}
	if !res2.IsPangram() {
		t.Fatal("Expected a pangram")
	}
}

// todo test sorting
