// repeatalpha
// Instructions

// Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...
// Expected Function

// func RepeatAlpha(s string) string {
// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.RepeatAlpha("abc"))
// 	fmt.Println(piscine.RepeatAlpha("Choumi."))
// 	fmt.Println(piscine.RepeatAlpha(""))
// 	fmt.Println(piscine.RepeatAlpha("abacadaba 01!"))
// }

// And its output:

// $ go run . | cat -e
// abbccc$
// CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
// $
// abbacccaddddabba 01!$
// $

package main

import (
	"os"
	"fmt"

	// "github.com/01-edu/z01"
)

func main(){
args := os.Args
if len(args) != 2 {
	return
}
str := args[1]
repeat := 0
// for _, ch := range str{
// 	if ch >= 'a' && ch <= 'z'{
// 		repeat = int(ch - 'a' + 1)
// 	}else if ch >= 'A'&& ch <= 'Z'{
// 		repeat = int(ch -'A' + 1)
// 	}else {
// 		repeat = 1
// 	}
// 	for i := 0; i < repeat; i++{
// 		z01.PrintRune(ch)
// 	}
// }
// z01.PrintRune('\n')
}

func Indexof(s string, ch rune)int{
	indexof:= 0
	for i, c := range s{
		if c == ch{
			indexof = i
		}else{
			indexof = -1
		}
	}
	return indexof
}
