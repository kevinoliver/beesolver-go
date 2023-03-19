package main

import "testing"

func TestNewPuzzleOtherLettersMustBeRightLength(t *testing.T) {
	if _, err := NewPuzzle('z', "1235"); err == nil {
		t.Error()
	}
	if _, err := NewPuzzle('z', "1234567"); err == nil {
		t.Error()
	}
}

func TestNewPuzzleOtherLettersCannotHaveDuplicates(t *testing.T) {
	if _, err := NewPuzzle('z', "abcdee"); err == nil {
		t.Error()
	}
}

func TestNewPuzzleOtherLettersCannotHaveRequiredLetter(t *testing.T) {
	if _, err := NewPuzzle('z', "abcdez"); err == nil {
		t.Error()
	}
}

func TestResultForValidWordsThatArentPangrams(t *testing.T) {
	p, _ := NewPuzzle('d', "ogselm")
	if res := p.ResultFor("dogs"); !res.IsValid() || res.IsPangram() {
		t.Error()
	}
	if res := p.ResultFor("doom"); !res.IsValid() || res.IsPangram() {
		t.Error()
	}
	if res := p.ResultFor("does"); !res.IsValid() || res.IsPangram() {
		t.Error()
	}
	if res := p.ResultFor("moods"); !res.IsValid() || res.IsPangram() {
		t.Error()
	}
}

func TestResultForValidWordsThatArePangrams(t *testing.T) {
	p, _ := NewPuzzle('d', "ogselm")
	if res := p.ResultFor("dogselm"); !res.IsValid() || !res.IsPangram() {
		t.Error()
	}
	if res := p.ResultFor("dogselmm"); !res.IsValid() || !res.IsPangram() {
		t.Error()
	}
	if res := p.ResultFor("dogselmdogselm"); !res.IsValid() || !res.IsPangram() {
		t.Error()
	}
}

func TestResultForWordsMissingRequiredLetter(t *testing.T) {
	p, _ := NewPuzzle('d', "ogselm")
	result := p.ResultFor("logs")
	if result.IsValid() {
		t.Error()
	}
	if result.IsPangram() {
		t.Error()
	}
}

func TestResultForWordsWithUnallowedLetters(t *testing.T) {
	p, _ := NewPuzzle('d', "ogselm")
	if p.ResultFor("deal").IsValid() {
		t.Error()
	}
	if p.ResultFor("seal").IsValid() {
		t.Error()
	}
}
