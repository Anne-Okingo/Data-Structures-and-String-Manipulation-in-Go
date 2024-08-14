// Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number
//  defined as a string into a number defined as an int.

// Atoi returns 0 if the string is not considered as a valid number.
// For this exercise only valid string will be tested. They will only contain one or several digits as characters.

// For this exercise the handling of the signs + or - does not have to be taken into account.

// This function will only have to return the int. For this exercise the error return of Atoi is not required.

package main

import "fmt"

func Atoi(str string) int {
	if len(str) <= 0 {
		return 0
	}

	result := 0
	for _, ch := range str {
		result = result*10 + int(ch -'0')
	}
	return result
}

func main() {
	fmt.Println(Atoi("1234"))
	fmt.Println(Atoi(""))
}
