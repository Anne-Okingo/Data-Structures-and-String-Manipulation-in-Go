// iscapitalized
// Instructions

// Write a function IsCapitalized() that takes a string as an argument and returns true if each word in the string begins with either an uppercase letter or a non-alphabetic character.

//     If any of the words begin with a lowercase letter return false.
//     If the string is empty return false.

// Expected function

// func IsCapitalized(s string) bool {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(IsCapitalized("Hello! How are you?"))
// 	fmt.Println(IsCapitalized("Hello How Are You"))
// 	fmt.Println(IsCapitalized("Whats 4this 100K?"))
// 	fmt.Println(IsCapitalized("Whatsthis4"))
// 	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
// 	fmt.Println(IsCapitalized(""))
// }

// And its output:

// $ go run .
// false
// true
// true
// true
// true
// false

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

// func IsCapitalized(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	} else {
// 		for i := 0; i < len(s); i++ {
// 			if i == 0 && (s[i] >= 'a' && s[i]<= 'z'){
// 				return false
// 			}
// 			if s[i] == ' ' && (s[i+1] >= 'a' && s[i+1]<= 'z') {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func IsCapitalized(s string) bool{
// 	if len(s) == 0{
// 		return false
// 	}

// 	for i := 0; i < len(s); i++{
// 		if i == 0 && (s[i] >= 'a' && s[i] <= 'z'){
// 			return false
// 		}
// 		if s[i] == ' ' && (s[i+1] >= 'a' && s[i+1] <= 'z') && i != len(s)-1{
// 			return false
// 		}
// 	}
// 	return true
// }

func IsCapitalized(s string) bool{
	if len(s) == 0{
		return false
	}
	for i, ch := range s{
		if (ch >= 'a' && ch <= 'z') && i==0 || s[i] == ' ' && (s[i+1] >= 'a' && s[i+1] <= 'z'){
			return false
		}
	}
	return true
}
