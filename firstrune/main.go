// Write a function that returns the first rune of a string.

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(FirstRune("Alice Anne"))
	z01.PrintRune(10)
}

func FirstRune(str string) rune {
	s := []rune(str)
	return s[0]
}



