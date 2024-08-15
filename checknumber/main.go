// checknumber
// Instructions

// Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.
// Expected function

// func CheckNumber(arg string) bool {
// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(CheckNumber("Hello"))
// 	fmt.Println(CheckNumber("Hello1"))
// }

// And its output:

// $ go run .
// false
// true
// $

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
func CheckNumber(str string)bool{
	for _, ch := range str{
		if ch >= '0' && ch <= '9'{
			return true
		}
	}
	return false
}