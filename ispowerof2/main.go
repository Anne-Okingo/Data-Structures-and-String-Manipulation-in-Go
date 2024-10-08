// ispowerof2
// Instructions

// Write a program that determines if a given number is a power of 2. A number n is a power of 2 if it meets the following condition: n = 2 m where m is another integer number.

// For example, 4 = 2 2 or 16 = 2 4 are power of 2 numbers while 24 is not.

// This program must print true if the given number is a power of 2, otherwise it has to print false.

//     If there is more than one or no argument, the program should print nothing.

//     When there is only one argument, it will always be a positive valid int.

// Usage :

// $ go run . 2 | cat -e
// true$
// $ go run . 64
// true
// $ go run . 513
// false
// $ go run .
// $ go run . 64 1024
// $

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		os.Exit(0)
	}
	str := args[1]

	num, _ := strconv.Atoi(str)

	fmt.Println(Power(num))
}

func Power(nb int) bool {
	if nb <= 0 {
		return false
	}
	// return nb > 0 && nb & (nb-1) == 0
	return (nb & (nb-1)== 0)
}
