// Write a function that returns the first prime number that is equal or superior to the int passed as parameter.

// The function must be optimized in order to avoid time-outs with the tester.

// (We consider that only positive numbers can be prime numbers)

package main

import (
	"fmt"

)


func main() {
	fmt.Println(FindPrevPrime(18))
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func IsPrime(nb int)bool{
	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++{
		if nb % i == 0{
			return false
		}
	}
	return true
}

func FindPrevPrime(num int)int{
	if IsPrime(num){
		return num
	}else{
		for i := num ; i > 0; i--{
			if IsPrime(i){
				return i
			}
		}
	}
	return num
}
