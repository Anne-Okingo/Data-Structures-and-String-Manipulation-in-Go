// concatslice
// Instructions

// Write a function ConcatSlice() that takes two slices of integers as arguments and returns the concatenation of the two slices.
// Expected function

// func ConcatSlice(slice1, slice2 []int) []int {

// }

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
	"strconv"
)

func main() {
	ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6})
	// fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	// fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}

func ConcatSlice(s1, s2 []int) {
	result1 := ""
for i := 0; i < len(s1); i++{
		result1 += strconv.Itoa(s1[i]) + ", "
}

result2 := ""
for j := 0; j < len(s2); j++{
	if j != len(s2)-1{
result2 += strconv.Itoa(s2[j]) + ","
	}else{
		result2 += strconv.Itoa(s2[j]) 
	}
}
str  := result1 + result2
slice := []string{}
for _, ch := range str{
		slice = append(slice,string(ch))
}
fmt.Println(slice)
}

