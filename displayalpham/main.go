// displayalpham
// Instructions
// Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

// Usage
// $ go run . | cat -e
// aBcDeFgHiJkLmNoPqRsTuVwXyZ$
// $

package main

import(
	"fmt"
	"unicode"
)

func main(){
	for c := 'a'; c <= 'z'; c++{
		if(c -'a' + 1)%2 == 0{
			fmt.Printf("%c", unicode.ToUpper(c))
		}else{
			fmt.Printf("%c", c)
		}
	}
}