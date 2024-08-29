package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	word := os.Args[1:]
	if len(args) != 3 {
		return
	}

	words := ""
	for _, ch := range word {
		words = words + string(ch)
	}
	result := ""
	seen := make(map[rune]bool)
	for _, chr := range words {
		if !seen[chr] {
			seen[chr] = true
			result = result + string(chr)
		}
	}
	fmt.Println(result)
}
