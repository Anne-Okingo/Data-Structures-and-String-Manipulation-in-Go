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
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// fmt.Println(IsPrime(3))
	// fmt.Println(Itoa(-254))
	// fmt.Println(Atoi("-254"))
	// fmt.Println(FPrime(28))
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	num := Atoi(str)
	fprime := FPrime(num)
	// fmt.Println(fprime)
	factorz := ""
	for _, ch := range fprime {
		factorz += Itoa(ch) + "*"
	}
	factorz = factorz[:len(factorz)-1]
	for _, c := range factorz {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
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

func Atoi(s string) int {
	m := 1
	result := 0
	for i, ch := range s {
		if i == 0 && ch == '-' {
			m = -1
			continue
		}
		if i == 0 && ch == '+' {
			continue
		}
		if ch < '0' || ch > '9' {
			continue
		}
		result = result*10 + int(ch-'0')
	}
	return m * result
}

func Itoa(n int) string {
	sign := ""
	result := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	if n == 0 {
		return "0"
	}
	for n > 0 {
		result = string(n%10+'0') + result
		n /= 10
	}
	return sign + result
}

func FPrime(n int) []int {
	factors := []int{}
	for i := 2; i <= n; {
		for n%i == 0 && IsPrime(i) {
			factors = append(factors, i)
			n = n / i
		}
		i++
	}
	return factors
}
