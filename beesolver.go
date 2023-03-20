package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	d1 := NewDictionaryFromPath("/home/kevin/dev/beesolver-go/beesolver.go")
	w1 := d1.Words()
	fmt.Println(len(w1))

	d2 := NewDictionary()
	w2 := d2.Words()
	fmt.Println(len(w2))
}
