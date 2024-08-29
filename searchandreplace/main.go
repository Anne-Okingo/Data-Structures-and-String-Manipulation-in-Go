// searchreplace
// Instructions

// Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

//     If the number of arguments is different from 3, the program displays nothing.

//     If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').

// Usage

// $ go run . "hella there" "a" "o"
// hello there
// $ go run . "hallo thara" "a" "e"
// hello there
// $ go run . "abcd" "z" "l"
// abcd
// $ go run . "something" "a" "o" "b" "c"
// $

// package main

// import "fmt"

// func Replace(str, str1, str2 string) string{
// 	result := ""
// 	for _, ch := range str {
// 		for _, chr := range str1{
// 			if ch == chr {
// 				result = result +str2
// 			}else {
// 				result = result + string(ch)
// 			}
// 		}
// 	}
// 	return result
// }

// func main(){
// 	fmt.Println(Replace("hello", "l", "y"))
// }

package main

import(
"os"
"github.com/01-edu/z01"
)

func main(){
	args := os.Args

	if len(args) < 4{
		return
	}

	str1 := args[1]
	str2 := args[2]
	str3 := args[3]

	result := ""

	for _, ch := range str1{
		for _, chr := range str2{
			if ch == chr {
				result = result + str3
			}else {
				result = result + string(ch)
			}
		}
	}
	for _, chrr := range result{
		z01.PrintRune(chrr)
	}
	z01.PrintRune('\n')
}

