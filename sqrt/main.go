package main                                

import(
	"fmt"
)

func Sqrt(nb int) int {
	for i := 0; i <= nb; i++{
		if i * i == nb {
			return i
		}
	}
	return 0
}
 func main(){
	fmt.Println(Sqrt(4))
 }

//  package piscine

//  func Sqrt(nb int)int {
// 	for i := 0; i <= nb; i++{
// 		if i*i == nb {
// 			return i
// 		}
// 	}
// 	return 0
//  }