// reversestrcap
// Instructions

// Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').

// If there are no argument, the program displays nothing.
// Usage

// $ go run . "First SMALL TesT" | cat -e
// firsT smalL tesT$
// $ go run . "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
// seconD tesT iS A littlE easieR$
// bewarE it'S noT harD wheN $
//  gO A dernieR 0123456789 foR thE roaD E$
// $ go run .
// $

package main

import (
	"fmt"
	"os"
	// "github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 1 {
		return
	}

	args := os.Args[1:]
	// res := ""

	for _, str := range args {
		fmt.Println(RevStrCap(str))
	}
	// fmt.Println(res)
}

func RevStrCap(s string) string {
	result := ""
	for i, c := range s {
		if i != len(s)-1 {
			if Up(c) && s[i+1] != ' ' {
				result += string(c + 32)
			} else if Low(c) && s[i+1] == ' ' {
				result += string(c - 32)
			} else {
				result += string(c)
			}
		} else {
			if i == len(s)-1 && Low(c) {
				result += string(c - 32)
			} else {
				result += string(c)
			}
		}
	}

	return result
}

func Low(s rune) bool {
	return s >= 'a' && s <= 'z'
}

func Up(s rune) bool {
	return s >= 'A' && s <= 'Z'
}
