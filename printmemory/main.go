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
	"github.com/01-edu/z01"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

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

func PrintMemory(arr [10]byte){
	result := []string{}
	for _, ch := range arr{
		if ch == 0{
			result = append(result, "00")
		}else{
			result = append(result, (Hex(int(ch))))
		}
	}
	for i := 0; i < len(result);i++{
		PrintS(result[i])
		if i == 3 || i == 7 || i == 9{
			z01.PrintRune('\n')
		}else if i < len(result) -1{
			z01.PrintRune(' ')
		}
	}

	for _, ch := range arr{
		if ch > 32 && ch <= 126{
			z01.PrintRune(rune(ch))
		}else{
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func PrintS(s string){
	for _, ch := range s{
		z01.PrintRune(ch)
	}

}

func Hex(n int)string{
	hex := "0123456789abcdef"
	result := ""

	if n  == 0{
		return "0"
	}

	for n > 0{
		remainder := n%16
		result = string(hex[remainder]) + result
		n /= 16
	}
	if len(result) == 1{
		result = "0" + result
	}
	return result
}
