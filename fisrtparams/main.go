// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	args := os.Args[1]

// 	for _, ch := range args {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }

package main

import(
"os"
"github.com/01-edu/z01"
)

func main(){
	arg := os.Args[1:]
	if len(arg) == 0{
		return
	}

	//args := arg[1]

	for _, ch := range arg[0] {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')

	// arg := os.Args[1:]
	// if len(arg) == 0{
	// 	return
	// }

	// //args := arg[1]
	// last := arg[len(arg)-1]

	// for _, ch := range last {
	// 	z01.PrintRune(ch)
	// }
	// z01.PrintRune('\n')
}
