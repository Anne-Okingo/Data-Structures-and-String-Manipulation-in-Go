// wordflip
// Instructions

// Write a function WordFlip() that takes a string as input and returns it in reverse order.

//     The output should be followed by a newline \n.
//     If the string is empty, return Invalid Output.
//     Ignore multiple spaces between words and trim any leading or trailing spaces in the string.

// Expected function

// func WordFlip(str string) string {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Print(piscine.WordFlip("First second last"))
// 	fmt.Print(piscine.WordFlip(""))
// 	fmt.Print(piscine.WordFlip("     "))
// 	fmt.Print(piscine.WordFlip(" hello  all  of  you! "))
// }

// And its output:

// $ go run . | cat -e
// last second First$
// Invalid Output$
// $
// you! of all hello$

 package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {
	if str == ""{
		return "Invalid Input-\n"
	}

	result := ""
	slice := []string{}
for _, ch := range str{
if ch != ' '{
	result = result + string(ch)
}
if ch == ' ' && result != ""{
	slice = append(slice,result)
	result = ""
}
}
if result != ""{
	slice = append(slice, result)
}

flip := ""
for i := len(slice)-1; i >= 0; i--{
flip = flip + slice[i] + " " 
}
return flip + "\n"
}


