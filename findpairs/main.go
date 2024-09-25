// Find Pairs
// Instructions

// Write a program that finds all pairs of elements in an integer array that sum up to a given target value. The program should output a list of pairs, each representing the indices of the elements that form the pair.

// In this exercise you must take in consideration the following:

//     Ensure it's possible to have positive or negative integers in the array.
//     Ensure each element is used only once in a pair, although the element can be repeated in different pairs.
//     Allow for multiple pairs to sum up to the target value.
//     The output messages should follow the one given in the examples bellow.
//     Return the message "No pairs found." when no pair is present.
//     Return the message "Invalid target sum." if the target is invalid.
//     Return the message "Invalid number: " if the number in the array is invalid.
//     For any input format that deviates from the specified format "[1, 2, 3, 4, 5]" "6", the program will return an "Invalid input." error message.

// Let's consider the input arr = [1, 2, 3, 4, 5] and the target sum targetSum = 6. When we run the program, the findPairs() function will search for pairs in the array that sum up to targetSum.

// Here is some example of outputs:

// $ go run . "[1, 2, 3, 4, 5]" "6"
// Pairs with sum 6: [[0 4] [1 3]]
// $ go run . "[-1, 2, -3, 4, -5]" "1"
// Pairs with sum 1: [[0 1] [2 3]]
// $ go run . "[1, 2, 3, 4, 5]" "10"
// No pairs found.
// $ go run . "[-1, -2, -3, -4, -5]" "-5"
// Pairs with sum -5: [[0 3] [1 2]]
// $ go run . "[1, 2, 3, 4, 20, -4, 5]" "2 5"
// Invalid target sum.
// $ go run . "[1, 2, 3, 4, 20, p, 5]" "5"
// Invalid number: p
// $ go run . "[1, 2, 3, 4" "5"
// Invalid input.
// $ go run . "1, 2, 3, 4" "5"
// Invalid input.
// $


package main

import(
"fmt"
)

func main(){
	fmt.Println(Paired([]int{1, 2, 3, 4, 5}, 6))
}

func Paired(slice []int, target int)[][]int{
	pairs :=[][]int{}
	mapped := make(map[int]int)

	for _, ch := range slice{
		mapped[ch]++
	}
	 for i, _ := range slice{
		if i < len(slice) && slice[i] + slice[i+1] == target{
			pairs = append(pairs,[]int{i,i+1})
		} 
	 }
	 return pairs
}/
