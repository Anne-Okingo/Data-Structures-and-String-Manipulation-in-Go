// Write a function that counts the characters of a string and that returns that count.
package main

import "fmt"

func Strlen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
	// return len([]rune(str))
}

func main() {
	fmt.Println(Strlen("Alice Anne"))
	fmt.Println(Strlen("how are you"))
}
