package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	dictionary, err := NewDictionaryFromReader(strings.NewReader("ogselmd\ndoom\ndogs\ndotty\nseem\n"))
	if err != nil {
		t.Error(err)
	}
	puzzle, err := NewPuzzle('d', "ogselm")
	if err != nil {
		t.Error(err)
	}
	solver := Solver{dictionary, puzzle}
	results := solver.Solve()
	if len(results) != 3 {
		t.Errorf("Expected 3 results, got %d", len(results))
	}
	if results[0].Word() != "dogs" {
		t.Errorf("Expected 'dogs' got '%s'", results[0].Word())
	}
	if results[1].Word() != "doom" {
		t.Errorf("Expected 'doom' got '%s'", results[1].Word())
	}
	res2 := results[2]
	if res2.Word() != "ogselmd" {
		t.Errorf("Expected 'ogselmd' got '%s'", res2.Word())
	}
	if !res2.IsPangram() {
		t.Error("Expected a pangram")
	}
}

// todo test sorting
