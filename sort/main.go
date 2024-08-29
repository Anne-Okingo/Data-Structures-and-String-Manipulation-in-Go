package main

import "fmt"

// func Sort(slice []int)[]int{
// 	n := len(slice)

// 	for i := 0 ; i < n-1; i++{
// 		for j := 0; j < n-i-1; j++{
// 			fmt.Print(j)
// 			if slice[j] > slice[j+1]{
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}

// return slice
// }

func Sort(nb []int)[]int{
	n := len(nb)

	for i := 0; i < n-1; i++{
		for j := 0; j < n-i-1; j++{
			if nb[j] > nb[j+1]{
				nb[j], nb[j+1] = nb[j+1], nb[j]		
				}
		}
	}
	return nb
}

func main(){
	fmt.Println(Sort([]int{3, 4, 2, 6, 5, 3,}))
}
