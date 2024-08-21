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
	if len(os.Args) != 3{
		return
	}
	s1, s2 := os.Args[1],os.Args[2]

	n1, n2 := Atoi(s1), Atoi(s2)

	for n2 != 0{
		n1,n2 = n2, n1%n2
	}
	fmt.Println(n1)

	// fmt.Println(Atoi("254"))

}

func Atoi(s string)int{
	multiplier := 1
	result := 0
	for i, ch := range s{
		if ch == '-' && i == 0{
			multiplier = -1
			s = s[1:]
			continue
		}
		if ch == '+' && i == 0{
			multiplier = 1
			s= s[1:]
			continue
		}
	}
	for  _, ch := range s{
		if ch < '0' || ch > '9'{
			continue
		}else{
			result = result * 10 + int(ch - '0')
		}
	}
	return  result * multiplier
}


