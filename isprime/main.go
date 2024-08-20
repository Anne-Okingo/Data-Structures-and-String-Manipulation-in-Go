// package main
 
// import(
// 	"fmt"
// )
// func main(){
// fmt.Println(IsPrime(4))
// }

// func IsPrime(nb int)bool{
// 	if nb < 1 {
// 		return false
// 	}else if nb == 2 {
// 		return true
// 	}
// 	for i := 0; i <= nb; i++{
// 		if nb % 2 == 0 {
// 			return false
// 		}
// 	}
// 	return true
// } 

// package main

// import "fmt"

// func IsPrime(nb int)bool {
// 	if nb <= 0 {
// 		return false
// 	}
// 	if nb == 2 {
// 		return true
// 	}
// 	for i := 0; i <= nb; i++{
// 		if nb % 2 == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main(){
// 	fmt.Println(IsPrime(17))
// 	fmt.Println(IsPrime(2))
// 	fmt.Println(IsPrime(12))
// 	fmt.Println(IsPrime(0))
// }
package main

import "fmt"

func IsPrime(num int)bool{
	if num <= 0{
		return false
	}

	// if num == 2 {
	// 	return true
	// }

	for i := 2; i <= num/2; i++{
		if num % i == 0 {
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(IsPrime(17))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(1287))
	fmt.Println(IsPrime(0))
}