// grouping
// Instructions
// Write a program that receives two strings and replicates the use of brackets in regular expressions. Brackets in regular expressions returns the words that contain the expression inside of it.

// The program should handle the "|" operator, which searches for both strings on each side of the operator.

// The output of the program should be, the results of the regular expression, numbered and displayed by the order of appearance in the string.

// If the number of arguments is different from 2, if the regular expression is not valid, if the last argument is empty or if there are no matches, the program should print nothing.

// Usage
// $ go run . "(a)" "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"
// 1: heavy
// 2: steady
// 3: heavy
// $ go run . "(e|n)" "I currently have 4 windows opened up… and I don’t know why."
// 1: currently
// 2: currently
// 3: have
// 4: windows
// 5: opened
// 6: opened
// 7: and
// 8: don’t
// 9: know
// $ go run . "(hi)" "He swore he just saw his sushi move."
// 1: his
// 2: sushi
// $ go run . "(s)" ""
// $ go run . "i" "Something in the air"
// $

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		return
	}
	str := args[1]

	if len(str) < 2 || str[0] != '(' || str[len(str)-1] != ')' {
		return
	}
	str = str[1 : len(str)-1]

	slice := Split(str, "|")

	sliced := Split(args[2], " ")
	count := 0
	// fmt.Println(sliced)

	// fmt.Println(slice)
	for i, str := range sliced {
		for _, str1 := range slice {
			if Contains(str, str1) {
				count++
				if str[len(str)-1] == ',' {
					str = str[:len(str)-1]
				}
				fmt.Printf("%v:%v\n", count, sliced[i])
			}
		}
	}

}

func Split(str string, sep string) []string {
	result := []string{}
	start := 0

	lensep := len(sep)

	for i := 0; i <= len(str)-lensep; i++ {
		if str[i:i+lensep] == sep {
			result = append(result, str[start:i])
			start = i + lensep
		}
	}
	if start <= len(str)-1 {
		result = append(result, str[start:])
	}
	return result
}

func Contains(str1 string, str2 string) bool {
	lenstr2 := len(str2)
	for i := 0; i <= len(str1)- lenstr2; i++ {
			if str1[i:i+lenstr2] == str2 {
				return true
		}
	}
	return false
}
