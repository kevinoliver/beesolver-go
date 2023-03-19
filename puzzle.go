package main

import (
	"fmt"
)

type Puzzle struct {
	requiredLetter rune
	allowedLetters map[rune]bool
}

func NewPuzzle(requiredLetter rune, otherLetters string) (Puzzle, error) {
	if len(otherLetters) != 6 {
		return Puzzle{}, fmt.Errorf("must have 6 otherLetters, found %d", len(otherLetters))
	}
	allowedLetters := make(map[rune]bool)
	for _, ch := range otherLetters {
		_, exists := allowedLetters[ch]
		if exists {
			return Puzzle{}, fmt.Errorf("otherLetters cannot have any duplicates, found: '%s'", string(ch))
		}
		allowedLetters[ch] = true
	}
	if _, exists := allowedLetters[requiredLetter]; exists {
		return Puzzle{}, fmt.Errorf("otherLetters cannot contain the requiredLetter: '%s'", string(requiredLetter))
	}
	allowedLetters[requiredLetter] = true
	return Puzzle{requiredLetter, allowedLetters}, nil
}

func (puzzle Puzzle) ResultFor(candidate string) Result {
	foundRequired := false
	allAreAllowed := true
	for _, candidateCh := range candidate {
		if candidateCh == puzzle.requiredLetter {
			foundRequired = true
		}
		if _, exists := puzzle.allowedLetters[candidateCh]; !exists {
			allAreAllowed = false
			break
		}
	}
	if foundRequired && allAreAllowed {
		uniqueLetters := make(map[rune]bool)
		for _, candidateCh := range candidate {
			uniqueLetters[candidateCh] = true
		}
		if len(uniqueLetters) == 7 {
			return NewPangramResult(candidate)
		} else {
			return NewValidResult(candidate)
		}
	} else {
		return NewInvalidResult(candidate)
	}
}
