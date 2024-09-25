// hashcode
// Instructions

// Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

//     The hash equation is computed as follows:

// (ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

//     If the resulting character is unprintable add 33 to it.

// Expected function

// func HashCode(dec string) string {
// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(HashCode("A"))
// 	fmt.Println(HashCode("AB"))
// 	fmt.Println(HashCode("BAC"))
// 	fmt.Println(HashCode("Hello World"))
// }

// And its output:

// $ go run .
// B
// CD
// EDF
// Spwwz+bz}wo

package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	mapped := make([]rune,len(dec))
	lend := len(dec)
	result := 0

	for i, ch := range dec{
		result = (int(ch) +lend) % 127

		if result < 32|| result > 126{
			result = result + 33
		}
		mapped[i] = rune(result)
	}
	return string(mapped)
}
// func HashCode(dec string) string {
// 	lend := len(dec)
// 	hashed := ""
// 	calc := 0
// 	for _, ch := range dec{
// 		calc = (int(ch) + lend) % 127
// 		if calc < 32 || calc > 126{
// 			calc += 33
// 		}
// 		hashed += string(calc)
// 	}
// return hashed
// }