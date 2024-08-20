// addprimesum
// Instructions

// Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

//     If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.

// Usage

// $ go run . 5
// 10
// $ go run . 7
// 17
// $ go run . -2
// 0
// $ go run . 0
// 0
// $ go run .
// 0
// $ go run . 5 7
// 0
// $

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	str := args[1]
	num := Atoi(str)
	sum := PrimeSum(num)
	result := Itoa(sum)
	for _, ch := range result{
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	multiplier := 1
	result := 0

	if s[0] == '-' {
		multiplier = -1
		s = s[1:]
	} else if s[0] == '+' {
		multiplier = 1
		s = s[1:]
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			continue
		} else {
			result = result*10 + int(ch-'0')
		}
	}
	return multiplier * result
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = n * -1
	}

	result := ""
	for n > 0 {
		result = string(n%10 + '0') + result
		n /= 10
	}
	return sign + result
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeSum(n int)int{
	sum := 0
	for i := 0;i < n; i++{
		if IsPrime(i){
			sum +=i
		}
	}
	return sum
}
