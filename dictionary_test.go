package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed dictionary/test-dictionary
var dictFile string

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func assertTrue(t *testing.T, value bool) {
	if !value {
		t.Error()
	}
}

func assertFalse(t *testing.T, value bool) {
	if value {
		t.Error()
	}
}

func TestNewDictionaryFromPathFiltersOutShortWords(t *testing.T) {
	words := NewDictionaryFromReader(strings.NewReader(dictFile)).Words()
	assertTrue(t, contains(words, "dogs"))
	assertFalse(t, contains(words, "cat"))
}

func TestNewDictionaryFromStringRemovesDuplicates(t *testing.T) {
	words := NewDictionaryFromReader(strings.NewReader("dogs\ndogs\ndogs\n")).Words()
	assertTrue(t, contains(words, "dogs"))
	if len(words) != 1 {
		t.Errorf("Expected 1 word, got %d. Words: %s", len(words), words)
	}
}

func TestNewDictionaryFromStringNormalizesAccents(t *testing.T) {
	words := NewDictionaryFromReader(strings.NewReader(dictFile)).Words()
	assertTrue(t, contains(words, "eclair"))
}
