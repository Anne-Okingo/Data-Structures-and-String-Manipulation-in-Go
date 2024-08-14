// Write a function that takes an int min and an int max as parameters.
// The function should return a slice of ints with all the values between the min and max.

// Min is included, and max is excluded.

// If min is greater than or equal to max, a nil slice is returned.

// make is not allowed for this exercise.

package main

import "fmt"

func AppendRange(min int, max int)[]int{
	if min >= max {
		return []int{}
	}

	result := []int{}
	for i := min; i < max; i++{
		result = append(result, i)
	}
	return result
}

 func main(){
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(4, 16))
	fmt.Println(AppendRange(16, 16))
 }