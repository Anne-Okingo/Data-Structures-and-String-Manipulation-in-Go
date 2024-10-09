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

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println(ParseInput(""))
	args := os.Args
	if len(args) != 3 {
		fmt.Println("invalid input")
		return
	}
	input := ParseInput(args[1])
	target, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid target sum")
		os.Exit(0)
	}

	pairs := Pairs(input, target)
	if len(pairs) > 0 {
		fmt.Printf("Pairs with sum % d: %v\n", target, pairs)
	} else {
		fmt.Println("No Pairs Found")
	}
}

func Sep(s, sep string) []string {
	lensep := len(sep)
	start := 0
	result := []string{}

	for i := 0; i < len(s)-lensep; i++ {
		if s[i:i+lensep] == sep {
			result = append(result, s[start:i])

			start = i + lensep
		}
	}
	if start <= len(s)-1 {
		result = append(result, s[start:])
	}
	// fmt.Println(result)
	return result

}

func Pairs(n []int, target int) [][]int {
	result := [][]int{}
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i]+n[j] == target {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func ParseInput(input string)[]int{
	result := []int{}
	if len(input) < 3 || input[0] != '[' || input[len(input)-1] != ']'{
		fmt.Println("Invalid input.")
		os.Exit(0)
	}else{
		input = input[1:len(input)-1]
	}
	// fmt.Println(input)
	input2 := Sep(input,", ")
	// fmt.Println(input2)
	for _, value := range input2{
		num,err := strconv.Atoi(string(value))
		if err != nil{
			fmt.Printf("Invalid number: %v\n", value )
			os.Exit(0)
		}
		result = append(result,num)
	}
	// fmt.Println(result)
	return result
}
