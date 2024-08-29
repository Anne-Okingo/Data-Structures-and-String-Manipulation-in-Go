// reduceint
// Instructions

// The function should have as parameters a slice of integers a []int and a function f func(int, int) int. For each element of the slice, it should apply the function f func(int, int) int, save the result and then print it.
// Expected function

// func ReduceInt(a []int, f func(int, int) int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }

// And its output :

// $ go run .
// 1000
// 502
// 250
// $

package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

// import "fmt"

// func ReduceInt(a[]int, f func(int, int) int){
// 	result := a[0]

// 	for i := 1; i < len(a); i++{
// 		result = f(result, a[i])
// 	}
// 	fmt.Println(result)
// }

// func ReduceInt(nb[]int, f func(int, int)int){
// 	one := nb[0]
// 	for i := 1; i < len(nb); i++{
// 		one = f(one, nb[i])
// 	}
// 	fmt.Println(one)
// }

// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }

func main(){
	mul := func(num int, as int)int{
		return num * as
	}
	sum := func(num int, as int)int{
		return num + as
	}
	div := func(num int, as int)int{
		return num/as
	}
	nb :=[]int{500,2}
	ReduceInt(nb,mul)
	ReduceInt(nb, sum)
	ReduceInt(nb,div)
}

func ReduceInt(nb []int, f func(int, int)int){
	result := nb[0]
	for i := 1; i < len(nb) ; i++{
		result = f(result, nb[i])
	}
	num := strconv.Itoa(result)
	 for _, ch := range num{
		z01.PrintRune(ch)
	 }
	 z01.PrintRune('\n')
}


