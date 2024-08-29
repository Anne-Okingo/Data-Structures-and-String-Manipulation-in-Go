// Instructions

// Write a program that takes a positive (or zero) number expressed in base 10, and that displays it in base 16 (with lowercase letters) followed by a newline ('\n').

//     If the number of arguments is different from 1, the program displays nothing.
//     Error cases have to be handled as shown in the example below.

// Usage

// $ go run . 10
// a
// $ go run . 255
// ff
// $ go run . 5156454
// 4eae66
// $ go run .
// $ go run . "123 132 1" | cat -e
// ERROR$
// $
package main

import (
"os"
// "fmt"
)

func main(){
	if len(os.Args) != 2{
		return
	}
	// fmt.Println(Atoi("-9"))
}

func Atoi(str string)int{
	multiplier := 1
	result := 0
	for i, ch := range str{
		if ch == '-' && i == 0{
			multiplier = -1
			str = str[1:]
			continue
		}else if ch == '+' && i == 0{
			multiplier = 1
			str = str[1:]
			continue
		}
		result = result * 10 + int(ch - '0')
	}
	return multiplier * result
}


package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	for _, ch := range args{
		if ch >= '0' && ch <= '9'{
			os.Stdout.WriteString("ERROR")
		}
	}

	num := Atoi(args)
	res := ""
	hexvals := "0123456789abcdef"

	for num > 0 {

		res = string(hexvals[num%16]) + res

		num /= 16
	}
	fmt.Println(res)
}
