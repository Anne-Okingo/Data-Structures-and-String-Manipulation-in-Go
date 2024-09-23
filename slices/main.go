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
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	elems := [][]interface{}{
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-3,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 4,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-2, -1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 0,
		},
	}

	s := random.StrSlice(chars.Words)

	elems = append(elems, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = random.StrSlice(chars.Words)
		elems = append(elems, []interface{}{s, random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)})
	}

	for _, a := range elems {
		challenge.Function("Slice", Slice, solutions.Slice, a...)
	}
}

func Slice(a []string, nbrs ...int) []string {
	start := 0
	end := len(a)

	if len(nbrs) >= 1 {
		start = nbrs[0]
		if start < 0 {
			start = len(a) + start
		}
		if start < 0 {
			start = 0
		}
		if start > len(a) {
			start = len(a)
		}
	}
	if len(nbrs) >= 2 {
		end = nbrs[1]

		if end < 0 {
			end = len(a) + end
		}
		if end < 0 {
			end = 0
		}
		if end > len(a) {
			end = len(a)
		}
		if start > end {
			return nil
		}
	}
	return a[start:end]
}
