// Write a program that displays the number of arguments passed to it. This number will be followed by a newline ('\n').

// If there is no argument, the program displays 0 followed by a newline.
// package main

// import (
// 	"os"
// 	"strconv"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	args := os.Args[1:]
// 	count := 0
// 	for _, cha := range args{
// 		for range cha{
// 			count++
// 		}
// 	}
// 	str := strconv.Itoa(count)
// 		for _,ch := range(str) {
// 			z01.PrintRune(ch)
// 		}
// 		z01.PrintRune('\n')

	// args := os.Args[1:]
	// count := 0
	//  for _, ch := range args {
	// 	for range ch {
	// 		count++
	// 	}
	// 	z01.PrintRune(rune(count + '0'))
	//  }
	//  z01.PrintRune('\n')

	// }
// package main

// import (
// "os"
// "github.com/01-edu/z01"
// )

// func main(){
// 	count := 0

// 	args := os.Args[1:]

// 	for range args{
// 		count++
// 	}

// 	str := ""
// 	if count < 0{
// 		str = "-1"
// 	} else {
// 		str = ""
// 	}

// 	result := ""
// 	for count != 0{
// 		mod := count % 10
// 		start := '0'
// 		for i := 0; i < mod; i++{
// 			start ++
// 		}
// 		result = string(start) + result
// 		count = count / 10
// 	}
// 	result = str + result

// 	for _, chr := range result{
// 		z01.PrintRune(chr)
// 	}
// 	z01.PrintRune('\n')
// }
/* Write a program that displays the number of arguments passed to it. This number will be followed by a newline ('\n').

If there is no argument, the program displays 0 followed by a newline. */
 package main

 import(
"os"
"github.com/01-edu/z01"
 )

 func main(){
	args := os.Args
	if len(args) == 1{
		z01.PrintRune('0')
	}
	// z01.PrintRune('\n')
	str := args[1:]

	count := 0
	for range str{
		count++
	}
	s := ""
	result := ""
	if count < 0{
		s = "-"
		count = -count
	}
	for count > 0{
		mod := count % 10
		start := '0'
		for i := 0; i < mod; i++{
			start++
		}
		result = result + string(start)
		count = count / 10
		result = result + s
	}
	for _, chr := range result{
		z01.PrintRune(chr)
	}
	z01.PrintRune('\n')
 }