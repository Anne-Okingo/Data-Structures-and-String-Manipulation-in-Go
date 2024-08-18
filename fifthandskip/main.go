// fifthandskip
// Instructions

// Write a function FifthAndSkip() that takes a string and returns another string. The function separates every five characters of the string with a space and removes the sixth one.

//     If there are spaces in the middle of a word, ignore them and get the first character after the spaces until you reach a length of 5.
//     If the string is less than 5 characters return Invalid Input followed by a newline \n.
//     If the string is empty return a newline \n.

// Expected function

// func FifthAndSkip(str string) string {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
// 	fmt.Print(piscine.FifthAndSkip("This is a short sentence"))
// 	fmt.Print(piscine.FifthAndSkip("1234"))
// }

// And its output:

// $ go run . | cat -e
// abcde ghijk mnopq stuwx z$
// Thisi ashor sente ce$
// Invalid Input$

package main

import (
	"fmt"
	// "os"
	// "github.com/01-edu/z01"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
	// FifthAndSkip("This is a short sentence")
}

func FifthAndSkip(str string) string {
	if str == ""{
		return "\n"
	}
	if len(str) < 5{
		return"Invalid Input" + "\n"
	}

	words := ""
	for _, ch := range str{
		if ch == ' '{
			continue
		}else{
			words += string(ch)
		}
	}
	// fmt.Println(words)
	result := ""
	for i := 0; i <= len(words)-1; i+=5{
		if i + 5 > len(words){
			result += string(words[i:])
		}else{
			result += words[i:i+5] + " "
		}
		i++
	}
	return result + "\n"

	}


