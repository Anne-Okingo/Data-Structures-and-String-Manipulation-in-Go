// weareunique
// Instructions

// Write a function that takes two strings's' and returns the number of characters that are not included in both, without repeating characters.

//     If there is no unique characters return 0.
//     If both strings are empty return -1.

// Expected function

// func WeAreUnique(str1 , str2 string) int {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(WeAreUnique("foo", "boo"))
// 	fmt.Println(WeAreUnique("", ""))
// 	fmt.Println(WeAreUnique("abc", "def"))
// }

// And its output:

// $ go run .
// 2
// -1
// 6

package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
	fmt.Println(WeAreUnique("abc", "def"))
}


func WeAreUnique(s1,s2 string) int{
	if s1 == ""&& s2 == ""{
		return -1
	}
	count := 0
	// // result := ""
	// unik1 := make(map[rune]bool)
	// unik2 := make(map[rune]bool)

	// for _, ch := range s1{
	// 	unik1[ch]= true
	// }
	// for _, chr := range s2{
	// 	unik2[chr]=true
	// }
	// // result1:= ""
	//  for _, c := range s1{
	// 	if !unik2[c]{
	// 		count++
	// 		// result1 += string(c)
	// 	}
	//  }
	//  for _, chrr := range s2{
	// 	if !unik1[chrr]{
	// 		count++
	// 		// result += string(chrr)
	// 	}
	//  }
	//  return count
	
	mapped := make(map[rune]int)

	// for _, ch := range s1{
	// 	 mapped[ch]++
	// }
	s3 := s1+s2

	for _, c := range s3{
		mapped[c]++
	}
	// fmt.Println(mapped)

	for _, chr := range s1{
		if mapped[chr]== 1{
			count++
		}
	}
	for _, chh := range s2{
		if mapped[chh] == 1{
			count++
		}
	}
	return count
	
}


