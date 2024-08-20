// Write a program that displays "Hello World!" followed by a newline ('\n').
// package main

// import "github.com/01-edu/z01"

//	func main(){
//		str := "Hello World!"
//		for _, ch := range str{
//			z01.PrintRune(ch)
//		}
//		z01.PrintRune('\n')
//	}
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	str := "Hello World!"

	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
