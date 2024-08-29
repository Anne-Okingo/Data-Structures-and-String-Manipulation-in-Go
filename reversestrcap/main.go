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
)

func main() {
	// if len(os.Args)!= 2{
	// 	return
	// }
	slice := os.Args[1:]
	result2 := ""
	for i := 0; i < len(slice); i++ {
		sliced := Tolower(string(slice[i]))
		// fmt.Println(sliced)

		for i, ch := range sliced {
			if ch >= 'a' && ch <= 'z' {
				if (i < len(sliced)-1 && sliced[i+1] == ' ') || i == len(sliced)-1 {
					result2 += string(sliced[i] - 32) + " "
				} else {
					result2 += string(sliced[i])
				}
			}
		}
	}
	fmt.Println(result2)
}

// fmt.Println(Tolower("First SMALL TesT"))

func Tolower(str string) string {
	result := ""
	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			result += string(ch + 32)
		} else {
			result += string(ch)
		}
	}
	return result
}
