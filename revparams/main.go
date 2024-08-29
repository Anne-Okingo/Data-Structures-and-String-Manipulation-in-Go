package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, ch := range args{
		for i := len(ch)-1 ; i >= 0; i--{
			z01.PrintRune(rune(ch[i]))
		}
		z01.PrintRune('\n')
	}
}