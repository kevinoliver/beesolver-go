package main

import (
	"bufio"
	_ "embed"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//go:embed dictionary/american-english-large
var defaultDictionary string

var normalizer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

func removeAccents(str string) (string, error) {
	s, _, err := transform.String(normalizer, str)
	if err != nil {
		return "", err
	}
	return s, err
}

type Dictionary struct {
	words []string
}

func NewDictionary() Dictionary {
	return NewDictionaryFromReader(strings.NewReader(defaultDictionary))
}

func NewDictionaryFromPath(path string) Dictionary {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return NewDictionaryFromReader(file)
}

func NewDictionaryFromReader(reader io.Reader) Dictionary {
	uniqueWords := make(map[string]bool)
	words := []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// - Spelling Bee words must be 4 letters of more
		// - removing accents allows us to match on words like "éclair"
		// - Using a Set removes any duplicates
		line := scanner.Text()
		line, err := removeAccents(line)
		if err != nil {
			log.Fatal(err)
		}
		if len(line) > 3 {
			_, exists := uniqueWords[line]
			if !exists {
				uniqueWords[line] = true
				words = append(words, line)
			}
		}
	}
	return Dictionary{words}
}

func (dict Dictionary) Words() []string {
	return dict.words
}
