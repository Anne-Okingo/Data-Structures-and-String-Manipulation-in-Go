// unmatch
// Instructions

// Write a function, Unmatch, that returns the element of the slice that does not have a correspondent pair.

//     If all the number have a correspondent pair, it should return -1.

// Expected function

// func Unmatch(a []int) int {

// }

// Usage

// Here is a possible program to test your function :
// func Unmatch(a []int) int {

// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := []int{1, 2, 3, 1, 2, 3, 4}
// 	unmatch := piscine.Unmatch(a)
// 	fmt.Println(unmatch)
// }

// And its output :

// $ go run .
// 4
// $

package main

import (
	"fmt"

)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}

// func Unmatch(a []int) int {
// mapped := make(map[int]int)

// for _,n := range a{
// 	mapped[n]++
// }

// num := 0
// for k, v := range mapped{
// 	if v != 2{
// 		num += k	
// 	}
// }
// return num
// }

func Unmatch(a []int) int {
	result := 0

	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a); j++ {
			if a[i] != a[j] {
				result = a[i]
			}
		}
	}
	return result
}



