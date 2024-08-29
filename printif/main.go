// printif
// Instructions

// Write a function that takes a string as an argument and returns the letter G followed by a newline \n if the argument length is more than 3, otherwise returns Invalid Input followed by a newline \n.

//     If it's an empty string return G followed by a newline \n.

// Expected function

// func PrintIf(str string) string {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(PrintIf("abcdefz"))
// 	fmt.Print(PrintIf("abc"))
// 	fmt.Print(PrintIf(""))
// 	fmt.Print(PrintIf("14"))
// }

// And its output:

// $ go run . | cat -e
// G$
// Invalid Output$
// G$
// Invalid Output$

package main

import (
	"fmt"
	// "github.com/01-edu/z01"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(str string) string {
// 	if len(str) == 0 || len(str)> 3{
// 		return "G" + "\n"
// 	//}
// 	// count := 0
// 	// for range str {
// 	// 	count++
// 	// }
// 	// if count >= 3 {
// 	// 	return "G"
// 	 } //else {
// 		return "Invalid Input" + "\n"
// 	 //}
count := 0
for range str{
count++
}
if str == ""{
	return "G\n"
}
if count >3{
	return "G\n"
}else{
	return "Invalid Input\n"
}
	}

