// package main

// import "fmt"

// func main() {
// 	fmt.Println(Rot13("hello ANNE you are 123"))
// }
// func Rot13(str string) string {
// 	result := ""
// 	new := []rune{}
// 	for _,ch := range str{
// 		if ch >= 'a' && ch <= 'z'{
// 			new = []rune{(ch - 'a' + 13) % 26 + 'a'}
// 			result = result + string(new)
// 		} else if ch >= 'A' && ch <= 'Z'{
// 			new = []rune{(ch - 'A' + 13) % 26 + 'A'}
// 			result = result + string(new)
// 		} else {
// 			result = result + string(ch)
// 		}
// 	}
// 	return string(result)
// }

// package main

// import (
// 	"os"
// 	"github.com/01-edu/z01"
// )

// func main(){
// 	args := os.Args[1]
// 	if len(args) < 2{
// 		return
// 	}
// 	result := ""
// 	for _, ch := range args{
// 		if ch >= 'a' && ch <= 'z'{
// 			result = result + string((ch - 'a' + 13) % 26 + 'a')
// 		}else if ch >= 'A' && ch <= 'Z' {
// 			result = result + string((ch - 'A' + 13) % 26 + 'A')
// 		}else {
// 			result = result + string(ch)
// 		}
// 	}
// 	for _, ch := range result {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }

// package main

// import (
// "os"
// "github.com/01-edu/z01"
// )

// func main(){
// 	args := os.Args[1]
// 	result := ""
// for _, chr := range args{
// 		if chr >= 'a' || chr <= 'z'{
// 			result = result + string((chr - 'a' + 13)% 26 + 'a')
// 		}else if chr >= 'A' || chr <= 'Z'{
// 			result = result + string((chr - 'A' + 13) % 26 + 'A')
// 		}else {
// 			result = result + string(chr)
// 		}
// 	}
// 	for _, ch := range result{
// 		z01.PrintRune(ch)
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
		result = result + string((ch -'a' + 13)% 26 + 'a')
	}else if ch >= 'A' || ch <= 'Z'{
		result = result + string((ch - 'A' + 13)% 26 + 'A')
	}else {
		result = result + string(ch)
	}
 }
 for _, chr := range result{
	z01.PrintRune(chr)
 }
z01.PrintRune('\n')
}

