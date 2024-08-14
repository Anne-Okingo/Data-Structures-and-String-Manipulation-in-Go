// Write a program called alphamirror that takes a string as argument and displays this string
// after replacing each alphabetical character with the opposite alphabetical
// character.

// The case of the letter stays the same, for example :

// 'a' becomes 'z', 'Z' becomes 'A'
// 'd' becomes 'w', 'M' becomes 'N'

// The final result will be followed by a newline.

// If the number of arguments is not 1, the program will display only a newline.
// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	args := os.Args
// 	arg := args[1]

// 	if len(os.Args) != 2 {
// 		z01.PrintRune('\n')
// 	}
// 	result := ""
// 	for _, ch := range arg {
// 		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
// 			if ch >= 'a' && ch <= 'z' {
// 				result = result + string('z'-(ch-'a'))
// 			} else if ch >= 'A' && ch <= 'Z' {
// 				result = result + string('Z'-(ch-'A'))
// 			}
// 		}
// 	}
// 	for _, chr := range result {
// 		z01.PrintRune(chr)
// 	}
// 	z01.PrintRune('\n')
// }

package main

import (
"os"
"github.com/01-edu/z01"
)

func main(){
	arg := os.Args[1:]

	if len(arg) == 0{
		return
	}

	args := arg[0]

	result := ""
	for _, ch := range args{
		if ch >= 'a' || ch <= 'z'{
			result = result + string('a' - ch + 'z')
		}else if ch >= 'A' || ch <= 'Z'{
			result = result + string('A'- ch + 'Z')
		}else {
			result = result + string(ch)
		}
	}
	for _, chr := range result{
		z01.PrintRune(chr)
	}
	z01.PrintRune('\n')
}

