// cleanstr
// Instructions
// Write a program that takes a string, and displays this string with exactly:

// one space between words.
// without spaces nor tabs at the beginning nor at the end.
// with the result followed by a newline ("\n").
// A "word" is defined as a part of a string delimited either by spaces/tabs, or by the start/end of the string.

// If the number of arguments is not 1, or if there are no words to display, the program displays a newline("\n").

// Usage :
// $ go run . "you see it's easy to display the same thing" | cat -e
// you see it's easy to display the same thing$
// $ go run . " only    it's  harder   "
// only it's harder$
// $ go run . " how funny" "Did you   hear Mathilde ?"

// $ go run . ""

// $
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	args := os.Args[1]

// 	trim := strings.TrimSpace(args)

// 	word := strings.Fields(trim)

// 	result := strings.Join(word, " ")

// 	fmt.Println(result)
// }

package main

import (
	// "os"
	"fmt"
	"os"
	// "github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println()
// 		return
// 	}
// 	//
// 	str := os.Args[1]
// 	// str := "   you see it's easy       to display the same thing      "
// 	s := ""
// 	answer := Split(str)
// 	for i := 0; i < len(answer); i++ {
// 		if i != len(answer)-1{
// 			s += answer[i] + " "
// 		} else {
// 		s += (answer[i])
// 	}
// 	}
// 	PrintS(s)
// 	z01.PrintRune('\n')
// }

// func Split(str string) []string {
// 	sliced := []string{}
// 	result := ""
// 	for i := 0; i < len(str); i++ {
// 		if str[i] != ' ' {
// 			result = result + string(str[i])
// 		} else if str[i] == ' ' && result != "" {
// 			sliced = append(sliced, result)
// 			result = ""
// 		}
// 	}
// 	if result != "" {
// 		sliced = append(sliced, result)
// 	}
// 	return sliced
// }

//	func PrintS(str string) {
//		for _, c := range str {
//			z01.PrintRune(c)
//		}
//	}
func main() {
	if len(os.Args[1]) == 0 || len(os.Args) != 2 {
		fmt.Println()
		return
	}
	str := os.Args[1]
	result := ""
	sliced := []string{}
	for _, ch := range str {
		if ch != ' ' {
			result = result + string(ch)
		}else if ch == ' ' && result != ""{
			sliced = append(sliced, result)
			result = ""
		}
	}
	if result != ""{
		sliced = append(sliced, result)
	}
	rev := ""
	for i := 0; i <= len(sliced) -1; i++{
		if i != len(sliced) -1 {
			rev = rev + sliced[i] + " "
		}else {
			rev= rev + sliced[i]
		}
	}
	fmt.Println(rev)
}
