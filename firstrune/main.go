// Write a function that returns the first rune of a string.
// package main

// import "fmt"

// func main() {
// 	fmt.Println(FirstRune("Alice"))
// }

// func FirstRune(str string) string {
// 	new := ""
// 	for range str {
// 		new = string(str[0])
// 	}
// 	return new
// }
// package main

// import "fmt"

// func main(){
// 	fmt.Println(FirstRune("Hello!"))
// }

// func FirstRune(str string)rune{
// 	s := []rune(str)
// 	return (s[0])
// }

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(FirstRune("Alice Anne"))
	z01.PrintRune(10)
}

func FirstRune(str string) rune {
	s := []rune(str)
	return (s[0])
}



