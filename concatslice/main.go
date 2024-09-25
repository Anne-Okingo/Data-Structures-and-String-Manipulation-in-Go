// concatslice
// Instructions

// Write a function ConcatSlice() that takes two slices of integers as arguments and returns the concatenation of the two slices.
// Expected function

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(piscine.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{}))
// }

// And its output:

// $ go run .
// [1 2 3 4 5 6]
// [4 5 6 7 8 9]
// [1 2 3]

package main

import (
	"fmt"
	// "strconv"
)

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
	fmt.Println(ConcatSlice([]int{}, []int{}))
	// fmt.Println(Atoi("254"))
}

func ConcatSlice(s1,s2 []int)[]int{
	if s1 == nil && s2 == nil{
		return nil
	}else{
		return append(s1,s2...)
	}
	// sliced := make([]int,len(s1) + len(s2))
	// copy(sliced,s1)
	// copy(sliced[len(s1):], s2)
	// return sliced
}

// func ConcatSlice(s1, s2 []int) []int{
// 	result1 := ""
// 	// slice := []string{}
// for i := 0; i < len(s1); i++{
// 		result1 += strconv.Itoa(s1[i])
// 	}

// result2 := ""
// for j := 0; j < len(s2); j++{
// 	result2 += strconv.Itoa(s2[j])
// }
//  str := result1 + result2

// num := []int{}
// for i := 0; i < len(str); i++{
// 	num = append(num, Atoi(string(str[i])))
// }
// return num

// }

// func Atoi(s string) int {
// 	multi := 1
// 	result := 0
// 	for i := range s {
// 		if s[0] == '-' {
// 			multi = -1
// 			s = s[i:]
// 			continue
// 		}
// 		if s[0] == '+' {
// 			multi = 1
// 			s = s[i:]
// 			continue
// 		}
// 	}
// 	for _, ch := range s {
// 		result = result*10 + int(ch-'0')
// 	}
// 	return multi * result
// }
