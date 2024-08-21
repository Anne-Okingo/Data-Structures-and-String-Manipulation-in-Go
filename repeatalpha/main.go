// Instructions

// Write a program called repeat_alpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

// The result must be followed by a newline ('\n').

// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

// If the number of arguments is different from 1, the program displays nothing.

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	args := os.Args[1]

// 	if len(os.Args) < 2 {
// 		//z01.PrintRune('')
// 		return
// 	}

// 	// if len(args) != 1 {
// 	// 	return
// 	// }
// 	repeat := 1
// 	for _, ch := range args {
// 		if ch >= 'a' || ch <= 'z' {
// 			repeat = int(ch - 'a' + 1)
// 		} else if ch >= 'A' || ch <= 'Z' {
// 			repeat = int(ch - 'A' + 1)
// 		}
// 		for j := 0; j <= repeat; j++ {
// 			z01.PrintRune(ch)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	arg := os.Args

// 	if len(arg) != 2 {
// 		return
// 	}

// 	args := arg[1]

// 	repeat := 0
// 	for _, ch := range args {
// 		if ch >= 'a' && ch <= 'z' {
// 			repeat = int(ch - 'a' + 1)
// 		} else if ch >= 'A' && ch <= 'z' {
// 			repeat = int(ch - 'A' + 1)
// 		}else{
// 			repeat = 1
// 		}
// 		for i := 0; i < repeat; i++ {
// 			z01.PrintRune(ch)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
args := os.Args
if len(args) != 2 {
	return
}
str := args[1]
repeat := 0
for _, ch := range str{
	if ch >= 'a' && ch <= 'z'{
		repeat = int(ch - 'a' + 1)
	}else if ch >= 'A'&& ch <= 'Z'{
		repeat = int(ch -'A' + 1)
	}else {
		repeat = 1
	}
	for i := 0; i < repeat; i++{
		z01.PrintRune(ch)
	}
}
z01.PrintRune('\n')
}
