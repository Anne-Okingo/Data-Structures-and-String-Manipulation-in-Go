// printmemory
// Instructions

// Write a function that takes (arr [10]byte), and displays the memory as in the example.

// After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

// The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.
// Expected function

// func PrintMemory(arr [10]byte) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import "piscine"

// func main() {
// 	piscine.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
// }

// And its output :

// $ go run . | cat -e
// 68 65 6c 6c$
// 6f 10 15 2a$
// 00 00$
// hello..*..$
// $
package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	res := []string{}
	for _, ch := range arr {
		if ch == 0 {
			res = append(res, "00")
		} else {
			res = append(res, Hex(int(ch)))
		}
	}
	for i := 0; i < len(arr); i++ {
		Print(res[i])
		if i == 3 || i == 7 || i == 9 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
	for _, ch := range arr {
		if ch >= 32 && ch <= 126 {
			fmt.Print(string(ch))
		} else {
			fmt.Print(".")
		}
	}
	z01.PrintRune('\n')
}

func Print(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table2)
}

func Hex(nb int) string {
	hex := "0123456789abcdef"

	if nb == 0 {
		return "0"
	}

	result := ""
	for nb > 0 {
		mod := nb % 16
		result = string(hex[mod]) + result
		nb = nb / 16
	}
	return result
}
