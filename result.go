package main

type Result struct {
	word string

	isValid bool

	isPangram bool
}

func NewInvalidResult(word string) Result {
	return Result{word, false, false}
}

func NewPangramResult(word string) Result {
	return Result{word, true, true}
}

func NewValidResult(word string) Result {
	return Result{word, true, false}
}

func (result Result) Word() string {
	return result.word
}

func (result Result) IsValid() bool {
	return result.isValid
}

func (result Result) IsPangram() bool {
	return result.isPangram
}
