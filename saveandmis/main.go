// saveandmiss
// Instructions

// Write a function called SaveAndMiss() that takes a string and an int as an argument. The function should move through the string in sets determined by the int, saving the first set, omitting the second, saving the third, and so on, in a 'save' and 'miss' fashion until the end of the string is reached. Return a string containing the saved characters.

//     If the int is 0 or a negative number return the original string.

// Expected function

// func SaveAndMiss(arg string, num int) string {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.SaveAndMiss("123456789", 3))
// 	fmt.Println(piscine.SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
// 	fmt.Println(piscine.SaveAndMiss("", 3))
// 	fmt.Println(piscine.SaveAndMiss("hello you all ! ", 0))
// 	fmt.Println(piscine.SaveAndMiss("what is your name?", 0))
// 	fmt.Println(piscine.SaveAndMiss("go Exercise Save and Miss", -5))
// }

// And its output:

// $ go run . | cat -e
// 123789$
// abcghimnostuz$
// $
// hello you all ! $
// what is your name?$
// go Exercise Save and Miss$

package main

import (

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	structs := []struct {
		str string
		itt int
	}{
		{"123456789", 1},
		{"e 5£ @ 8* 7 =56 ;", 2},
		{"QKplq%QSw", 3},
		{"", 4},
		{"hello \\! n4ght cr3a8ure7 ", 5},
		{"Kimetsu no Yaiba", 6},
		{"8595485-52", 7},
		{"abcdefghijklmnopqrstuvwyz", 8},
		{"w58tw7474abc", 9},
		{"Po65 4o", 10},
	}

	for _, v := range structs {
		challenge.Function("SaveAndMiss",SaveAndMiss, solutions.SaveAndMiss, v.str, v.itt)
	}
}

func SaveAndMiss(arg string, num int) string {
	if num < 1{
		return arg
	}
	result := ""
for i := 0; i < len(arg); i++{
	if (i / num) % 2 == 0{
		result += string(arg[i])
	}
}
return result
	}