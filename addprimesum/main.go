// // addprimesum
// // Instructions

// // Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

// //     If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.

// // Usage

// // $ go run . 5
// // 10
// // $ go run . 7
// // 17
// // $ go run . -2
// // 0
// // $ go run . 0
// // 0
// // $ go run .
// // 0
// // $ go run . 5 7
// // 0
// // $

// package main

// import (
// 	 "fmt"
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {

// 	args := os.Args
// 	if len(args) != 2 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	num := Atoi(args[1])
// 	if num < 0{
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	// fmt.Println(num)
// 	result := PrimeSum(num)

// 	results := Itoa(result)
// // fmt.Println(results)
// 	for _, v := range results {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

// func IsPrime(n int) bool {
// 	if n < 2 {
// 		return false
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func PrimeSum(n int) int {
// 	sum := 0

// 	for i := 0; i <= n; i++ {
// 		if IsPrime(i) {
// 			sum += i
// 		}
// 	}
// 	return sum
// }

// func Atoi(s string) int {

// 	multiplier := 1

// 	for i, ch := range s {
// 		if ch == '-' && i == 0 {
// 			multiplier = -1
// 			s= s[1:]
// 		}
// 		if ch == '+' && i == 0 {
// 			multiplier = 1
// 			s= s[1:]
// 		}
// 	}
// 	result := 0

// 	for _, v := range s {
// 		// if !(v >= '0' && v <= '9') {
// 		// 	z01.PrintRune('0' + '\n')
// 		// 	os.Exit(0)
// 		// }
// 		result = result*10 + int(v-'0')
// 	}
// 	return result * multiplier
// }

// func Itoa(n int) string {
// 	sign := ""
// 	result := ""
// 	if n < 0 {
// 		sign = "-"
// 		n = -n
// 	}

// 	for n != 0 {
// 		startrne := '0'
// 		mod := n % 10
// 		for i := 0; i < mod; i++ {
// 			startrne++
// 		}
// 		result = string(startrne) + result
// 		n /= 10
// 	}

// 	return sign + result
// }



package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	// Check number of arguments first
	if len(args) != 1 {
		fmt.Println("0")
		return
	}

	num, err := strconv.Atoi(args[0])
	if err != nil || num < 0 {
		fmt.Println("0")
		return
	}

	results := PrimeSum(num)
	fmt.Println(results)
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeSum(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		if IsPrime(i) {
			sum += i
		}
	}
	return sum
}

