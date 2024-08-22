// Instructions

// Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

// The string will be followed by a newline ('\n').

// A word, in this exercise, is a sequence of visible characters.

// If the number of arguments is not 1, or if there are no word, the program displays nothing.
// Usage

// $ go run . "you   see   it's   easy   to   display   the   same   thing" | cat -e
// you   see   it's   easy   to   display   the   same   thing$
// $ go run . "   only  it's harder   " | cat -e
// only   it's   harder$
// $ go run . " how funny it is" "did you  hear, Mathilde ?" | cat -e
// $ go run .
// $

package main

import(
"os"
"fmt"
)

func main(){
	if len(os.Args) != 2{
		return
	}
	str := os.Args[1]

	result := ""
	sliced := []string{}

	for _, ch := range str{
		if ch != ' '{
			result += string(ch)
		}else if result != ""{
			sliced = append(sliced, result)
			result = ""
		}
	}
	if result != ""{
		sliced = append(sliced, result)
	}
	fmt.Println(sliced)
	expanded := ""
	for i, word := range sliced{
		if i != len(sliced)-1{
			expanded += word + "   "
		}else{
			expanded += word
		}
	}
	fmt.Println(expanded)
}