// revconcatalternate
// Instructions

// Write a function RevConcatAlternate() that receives two slices of int as arguments and returns a new slice with alternated values of each slice in reverse order.

//     The input slices can have different lengths.
//     The new slice should start with the elements from the largest slice first and when they became equal size slices, it should add an element of the first given slice.
//     If the slices are of equal length, the new slice should start with an element of the first slice.

//     Note: you can check the examples bellow for more details.

// Expected function

// func RevConcatAlternate(slice1,slice2 []int) []int {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
// fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
// 	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{}))
// }

// And its output:

// $ go run .
// [3 6 2 5 1 4]
// [9 8 7 3 6 2 5 1 4]
// [8 9 3 2 5 1 4]
// [3 2 1]
package main

import "fmt"

// func RevConcatAlternate(slice1, slice2 []int) []int {
// 	result := []int{}
// 	if len(slice1)== len(slice2){
// 		for i := len(slice1)-1; i >= 0; i--{
// 			result = append(result,slice1[i])
// 			result = append(result,slice2[i])
// 		}
// 	}
// 	if len(slice2) == 0{
// 		for i := len(slice1)-1; i >= 0; i--{
// 			result = append(result,slice1[i])
// 		}
// 	}
// 	if len(slice1) == 0{
// 		for i := len(slice2)-1; i >= 0; i--{
// 			result = append(result,slice2[i])
// 		}
// 	}
// 	if len(slice1) > len(slice2){
// 		for i := len(slice1)-1; i >= 0; i--{
// 			result = append(result, slice1[i])
// 			result = append(result,slice2[i])
// 		}
// 		result = append(result, slice1[len(slice2):]...)
// 	}
// 	return result

// }

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
func RevConcatAlternate(s1,s2 []int)[]int{
sliced := []int{}
len1, len2 := len(s1), len(s2)
i, j := len1-1, len2 -1

for i >= 0 && j >= 0{
	if len1 > len2{
		sliced = append(sliced,s1[i])
		i--
		len1--
	}
	if len2 > len1{
		sliced = append(sliced,s2[j])
		j--
		len2--
	}else{
		sliced = append(sliced,s1[i])
		i--
		len1--
	}
	if len1 <= 0{
		sliced = append(sliced,s2[j])
		j--
	}
	if len2 <= 0{
		sliced =append(sliced,s1[i])
		i--
	}
}
return sliced
}
