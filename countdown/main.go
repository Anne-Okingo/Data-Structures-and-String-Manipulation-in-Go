// Write a program that displays all digits in descending order, followed by a
// newline.
// package main 

// import (
// "github.com/01-edu/z01"
// )

// func main(){
// 	for i := '9'; i >= '0'; i--{
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }

package main

import "github.com/01-edu/z01"

func main(){
	for i := '9'; i >= '0'; i--{
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
