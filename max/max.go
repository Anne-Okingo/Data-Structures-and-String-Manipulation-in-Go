// Write a function Max that will return the maximum value in a slice of integers. If the slice is empty it will return 0.******sort*****
// package main

// import (
// 	"fmt"
// )

// func Max(a []int) int {
// 	// check if a is an empty slice of intigers then returns 0
// 	if len(a)== 0 {
// 		return 0
// 	}
// 	result := a[0]

// 	for i:= 1; i < len(a); i++ {
// 		if a[i] > result {
// 			result = a[i]
// 		}
// 	}
// 	return result
// }
// func main(){
// 	intg := []int{1, 2, 3, 4, 5, 6}
// 	max := Max(intg)
// 	fmt.Println(max)
// }

// package main

// import "fmt"

// func Max(nb []int) int {
// 	if nb == nil {
// 		return 0
// 	}

// 	for i := 1; i < len(nb); i++ {
// 		if nb[i] > nb[0] {
// 			nb[0] = nb[i]
// 		}
// 	}
// 	return nb[0]
// }

// func main() {
// 	fmt.Println(Max([]int{1, 2, 3, 4, 5, 6}))
// }

// package main

// import (
// 	"fmt"
// )

// func Max(nb []int) int {
// 	// if len(nb) == 0 {
// 	// 	return 0
// 	// }
// 	// for i := 0; i < len(nb); i++ {
// 	// 	if nb[i] > nb[0] {
// 	// 		nb[0] = nb[i]
// 	// 	}
// 	// }
// 	// return nb[0]

// 	if len(nb) == 0{
// 		return 0
// 	}
// for i := 0; i < len(nb)-1; i++{
// 	for j := 0; j < len(nb)-i-1; j++{
// 		if nb[j] > nb[j+1]{
// 			nb[j], nb[j+1]= nb[j + 1], nb[j]
// 		}
// 	}
// }
// return nb[len(nb)-1]

// }
package main

import (
"fmt"
)

// func Max(nb []int)int{
// for i := 0; i < len(nb)-1; i++{
// 	for j := 0; j < len(nb)-i-1;j++{
// 		if nb[j] > nb[j+1]{
// 			nb[j], nb[j+1] = nb[j+1],nb[j]
// 		}
// 	}
// }
// return nb[len(nb)-1]
// }

func Max(num[]int)int{
	for i := 0; i < len(num)-1; i++{
		for j := 0; j < len(num)-i-1; j++{
			if num[j] > num[j+1]{
				num[j], num[j+1] = num[j+1],num[j]
			}
		}
	}
	return num[len(num)-1]
}

func main() {
	fmt.Println(Max([]int{268, 89, 20, 6767, 8}))
	fmt.Println(Max([]int{23, 123, 1, 11, 55, 93}))
	fmt.Println(Max([]int{1, 3, 4, 5, 6, 7, 8}))
}
