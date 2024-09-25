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

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func randomSize() []int {
	var randSlice []int
	for i := 0; i <= random.IntBetween(0, 20); i++ {
		randSlice = append(randSlice, random.Int())
	}
	return randSlice
}

func main() {
	type node struct {
		slice []int
		ch    int
	}
	args := []node{
		{
			slice: []int{},
			ch:    0,
		}, {
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8},
			ch:    0,
		},
	}

	for i := 0; i <= 7; i++ {
		value := node{
			slice: randomSize(),
			ch:    random.IntBetween(0, 10),
		}
		args = append(args, value)
	}
	for _, arg := range args {
		challenge.Function("Chunk", Chunk, solutions.Chunk, arg.slice, arg.ch)
	}
}

func Chunk(slice []int, size int) {
	chunked := [][]int{}

	if size == 0{
		fmt.Println()
		return
	}

	for i := 0; i < len(slice); i += size{
		end := i + size

		if end > len(slice){
			end = len(slice)
		}
		chunked = append(chunked,slice[i:end])
	}
	fmt.Println(chunked)
}

// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }
