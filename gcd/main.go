// gcd
// Instructions

// Write a program that takes two string representing two strictly positive integers that fit in an int.

// The program displays their greatest common divisor followed by a newline ('\n').

// If the number of arguments is different from 2, the program displays nothing.

// All arguments tested will be positive int values.
// Usage

// $ go run . 42 10 | cat -e
// 2$
// $ go run . 42 12
// 6
// $ go run . 14 77
// 7
// $ go run . 17 3
// 1
// $ go run .
// $ go run . 50 12 4
// $

package main

import(
"os"
"fmt"
)

func main(){
	args := os.Args

	args1 := Atoi(args[1])
	args2 := Atoi(args[2])

	for args2 != 0{
		args1 , args2 = args2 , args1 % args2
	}
	fmt.Println(args1)

}

func Atoi(str string)int{
	multiplier := 1
	for i := 0; i < len(str); i++{
		if str[i] == '-'{
			multiplier = -1
			str = str[i:]
		}else if str[i] == '+'{
			multiplier = 1
			str = str[i:]
		}
	}

	result := 0

	for _, ch := range str{
		result = result * 10 + int(ch -'0')
	}
	return result * multiplier
}

