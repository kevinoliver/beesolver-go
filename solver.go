package main

import "sort"

type Solver struct {
	dictionary Dictionary
	puzzle     Puzzle
}

func (solver Solver) Solve() []Result {
	results := make([]Result, 0)
	for _, word := range solver.dictionary.Words() {
		result := solver.puzzle.ResultFor(word)
		if result.isValid {
			results = append(results, result)
		}
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Word() < results[j].Word()
	})
	return results
}
