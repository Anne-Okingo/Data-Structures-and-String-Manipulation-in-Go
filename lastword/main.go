// Write a program that takes a string and displays its last word, followed by a newline ('\n').

//     A word is a section of string delimited by spaces or by the start/end of the string.

//     The output will be followed by a newline ('\n').

//     If the number of arguments is different from 1, or if there are no word, the program displays nothing.

// package main

// import (
// 	"os"
// 	"github.com/01-edu/z01"
// )
// func main(){
// 	args := os.Args

// 	if len(args) != 2 {
// 		return
// 	}

// 	arg := args[1]
// 	check := false

// 	for _, c := range arg{
// 		if c != ' ' {
// 			check = true
// 			break
// 		}
// 	}
// 	if check {
// 		result := ""
// 	for i := len(arg)-1; i >= 0; i--{
// 		if string(arg[i]) != " "{
// 			result = string(arg[i]) + result
// 		}else if result != ""{
// 			break
// 		}
// 	}
// 	for _, chrr := range result{
// 		z01.PrintRune(chrr)
// 	}
// 	z01.PrintRune('\n')
// 	}
// }

// package main

// import (
// 	// "os"
// 	"fmt"
// 	"github.com/01-edu/z01"
// )

// func main() {
	// args := os.Args
	// if len(args) != 2 {
	// 	return
	// }
	// str := args[1]
	// check := false

	// for _, c := range str {
	// 	if c != ' ' {
	// 		check = true
	// 		break
	// 	}
	// }
	// if check {
	// 	result := ""

	// 	for i := len(str)-1; i > 0; i--{
	// 		if string(str[i]) != " "{
	// 			result = string(str[i]) + result
	// 		}else if result != ""{
	// 			break
	// 		}
	// 	}
	// 	for _, chh := range result{
	// 		z01.PrintRune(chh)
	// 	}
	// 	z01.PrintRune('\n')
	// }
	

// }

package main

import (
	"fmt"
	// "piscine"
	// "github.com/01-edu/z01"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not           "))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}


func LastWord(str string)string{
check := false
for _, ch := range str{
	if ch != ' '{
		check = true
		break
	}
}
result := ""
if check{
	for i := 0; i < len(str)-1; i--{
		if str[i] != ' '{
			result += string(str[i])
		}else if result != ""{
			break
		}
	}
}
return result + "\n"
}
