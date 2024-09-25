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
	fmt.Print(FifthAndSkip("abcd efghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))

	fmt.Print(FifthAndSkip("e5£jhy@a*7=56;"))
}

func FifthAndSkip(str string) string {
if str == ""{
	return "\n"
}
if len(str) < 5{
	return "Invalid Input" + "\n"
}
words := ""
for _, ch := range str{
	if ch != ' '{
		words += string(ch)
	}
}
result := ""
for i := 0; i < len(words);i+=5{
	start := i+5
	if start > len(words){
		result += string(words[i:])
	}else{
		result += string(words[i:start]) + " "
	}
	i++
}
return result + "\n"
	}


