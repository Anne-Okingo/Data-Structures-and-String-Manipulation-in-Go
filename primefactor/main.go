
// / README.md
// fprime
// Instructions

// Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

//     Factors must be displayed in ascending order and separated by *.

//     If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

// Usage

// $ go run . 225225
// 3*3*5*5*7*11*13
// $ go run . 8333325
// 3*3*5*5*7*11*13*37
// $ go run . 9539
// 9539
// $ go run . 804577
// 804577
// $ go run . 42
// 2*3*7
// $ go run . a
// $ go run . 0
// $ go run . 1
// $
package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsPrime(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Fprime(n int) []int {
	res := []int{}
	for i := 2; i <= n; {
		if IsPrime(i) && n%i == 0 {
			res = append(res, i)
			n = n / i
		} else {
			i++
		}
	}
	return res
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	num, err := strconv.Atoi(args)
	if err != nil || num <= 0 {
		return
	}

	factors := Fprime(num)
	// fmt.Println(factors)
	res := ""
	for _, c := range factors {
		res += strconv.Itoa(c) + "*"
	}
	res = res[:len(res)-1]
	fmt.Println(res)
}
