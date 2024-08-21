// slice
// Instructions

// The function receives a slice of strings and one or more integers, and returns a slice of strings. The returned slice is part of the received one but cut from the position indicated in the first int, until the position indicated by the second int.

// In case there only exists one int, the resulting slice begins in the position indicated by the int and ends at the end of the received slice.

// The integers can be negative.
// Expected function

// func Slice(a []string, nbrs... int) []string{

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
//     "fmt"
//     "piscine"
// )

// func main(){
//     a := []string{"coding", "algorithm", "ascii", "package", "golang"}
//     fmt.Printf("%#v\n", piscine.Slice(a, 1))
//     fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
//     fmt.Printf("%#v\n", piscine.Slice(a, -3))
//     fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
//     fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))
// }

// $ go run .
// []string{"algorithm", "ascii", "package", "golang"}
// []string{"ascii", "package"}
// []string{"ascii", "package", "golang"}
// []string{"package"}
// []string(nil)

package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
	fmt.Printf("%#v\n", Slice(a, 2, 1, 5, 7))
}

func Slice(a []string, nbrs ...int) []string {
	// // Handle the case when no numbers are provided
	// if len(nbrs) == 0 {
	//     return nil
	// }

	// start := nbrs[0]
	// end := len(a) // Default end position is the end of the slice

	// // Handle negative indexing
	// if start < 0 {
	//     start = len(a) + start
	// }

	// // If a second number is provided, use it as the end position
	// if len(nbrs) > 1 {
	//     end = nbrs[len(nbrs)-1]
	//     if end < 0 {
	//         end = len(a) + end
	//     }
	// }

	// // Handle cases where start or end is out of bounds
	// if start < 0 || start >= len(a) || start > end {
	//     return nil
	// }

	// if end > len(a) {
	//     end = len(a)
	// }

	// return a[start:end]

	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	end := len(a)

	if start < 0 {
		start = len(a) + start
	}

	if len(nbrs) > 1 {
		// start = nbrs[0]
		end = nbrs[len(nbrs)-1]

		if start < 0 {
			start = len(a) + start
		}

		if end < 0 {
			end = len(a) + end
		}

		if end > len(a) {
			end = len(a)
		}
		if start < 0 || start > end || start >= len(a) {
			return nil
		}
	}
	return a[start:end]
}
