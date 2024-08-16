// package main  

// import (
// "os"
// "strconv"
// "github.com/01-edu/z01"
// )

// func main(){
// 	args := os.Args[1:]
// 	count := 0
// 	for _, ch := range args {
// 		for range ch{
// 			count++
// 		}
// 	}
// 	str := strconv.Itoa(count)
// 	for _, chr := range str {
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
	args := os.Args[1:]

	count := 0
	for _, chrs := range args{
		for range(chrs) {
			count++
		}
	}

	str := ""
	for count != 0 {
		mod := count % 10
		startrune := '0'
		for i := 0; i < mod ; i++{
			startrune++
		}
		str = str + string(startrune)
		count = count/10
	}
	for j := len(str)- 1; j >= 0; j--{
		z01.PrintRune(rune(str[j]))
	}
	z01.PrintRune('\n')
}