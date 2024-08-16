// chunk
// Instructions

// Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

//     If the size is 0 it should print a newline ('\n').

// Expected function

// func Chunk(slice []int, size int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }

// And its output :

// $ go run .
// []

// [[0 1 2] [3 4 5] [6 7]]
// [[0 1 2 3 4] [5 6 7]]
// [[0 1 2 3] [4 5 6 7]]
// $


package main

import (
	"fmt"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
	Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
}
func Chunk(slice []int, size int){
	if size == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println([][]int{})
		return
	}
	chunked := [][]int{}
	for i := 0; i < len(slice); i = i + size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunked = append(chunked, slice[i:end])
	}
	fmt.Println(chunked) 
}
