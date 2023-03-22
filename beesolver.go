package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	var dictionaryPath = flag.String("dict", "", "Path to a custom dictionary")
	var wordsOutput = flag.Bool("words-output", true, "Default on. When off, the solution's words are hidden")
	var help = flag.Bool("help", false, "Print the usage and exit")
	flag.Usage = beeUsage
	flag.Parse()

	// if the help flag was passed in, handle that first
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "help" {
			usageAndExit(nil)
		}
	})
	if *help {
		usageAndExit(nil)
	}

	if err := validate(*dictionaryPath); err != nil {
		usageAndExit(err)
	}
	required := rune(flag.Arg(0)[0])
	others := flag.Arg(1)
	puzzle, err := NewPuzzle(required, others)
	if err != nil {
		usageAndExit(err)
	}
	var dictionary Dictionary
	var dictionaryName string
	if *dictionaryPath == "" {
		dictionary, err = NewDictionary()
		dictionaryName = "(default)"
	} else {
		dictionary, err = NewDictionaryFromPath(*dictionaryPath)
		dictionaryName = *dictionaryPath
	}
	if err != nil {
		usageAndExit(err)
	}

	fmt.Println("🐝")
	fmt.Println("Hello and welcome to Spelling Bee Solver")
	fmt.Println("🐝🐝")

	fmt.Println("🐝🐝🐝")
	fmt.Println("Required Letter:  ", string(required))
	fmt.Println("Other Letters:    ", others)
	fmt.Println("Dictionary:       ", dictionaryName)
	fmt.Println("Dictionary words: ", len(dictionary.Words()))
	fmt.Println("Solving now")
	fmt.Println("🐝🐝🐝🐝")

	solver := Solver{dictionary, puzzle}
	solutions := solver.Solve()

	fmt.Println("🐝🐝🐝🐝🐝")
	fmt.Println("Solved!")
	fmt.Println()
	fmt.Println("  Words: ", len(solutions))
	numPangrams := 0
	for _, solution := range solutions {
		if solution.IsPangram() {
			numPangrams++
		}
	}
	fmt.Println("  Pangrams: ", numPangrams)
	fmt.Println("🐝🐝🐝🐝🐝🐝")

	if *wordsOutput {
		for _, solution := range solutions {
			if solution.IsPangram() {
				fmt.Println(solution.Word(), " 🍳")
			} else {
				fmt.Println(solution.Word())
			}
		}
	}
}

func validate(dictionaryPath string) error {
	if flag.NArg() != 2 {
		return fmt.Errorf("requires 2 arguments: REQUIRED_LETTER OTHER_LETTERS. Got %d", flag.NArg())
	}
	if len(flag.Arg(0)) != 1 {
		return fmt.Errorf("REQUIRED_LETTER must be a single character. Got '%s'", flag.Arg(0))
	}
	requiredLetter := rune(flag.Arg(0)[0])
	otherLetters := flag.Arg(1)
	if len(otherLetters) != 6 {
		return fmt.Errorf("OTHER_LETTERS must be exactly 6 letters, found: %d", len(otherLetters))
	}
	letters := make(map[rune]bool)
	for _, ch := range otherLetters {
		_, exists := letters[ch]
		if exists {
			return fmt.Errorf("OTHER_LETTERS cannot have any duplicates, found: '%s'", string(ch))
		}
		letters[ch] = true
	}
	if _, exists := letters[requiredLetter]; exists {
		return fmt.Errorf("OTHER_LETTERS cannot contain the REQUIRED_LETTER: '%s'", string(requiredLetter))
	}
	if dictionaryPath != "" {
		if _, err := os.Stat(dictionaryPath); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("dictionary file does not exist: %s", dictionaryPath)
		}
	}

	return nil
}

func beeUsage() {
	fmt.Printf("Usage: beesolver REQUIRED_LETTER OTHER_LETTERS\n")
	flag.PrintDefaults()
}

func usageAndExit(err error) {
	flag.Usage()

	if err != nil {
		fmt.Println()
		fmt.Println(err.Error())
	}
	os.Exit(1)
}
