// Write a function that only counts the characters of a string and that returns that count.
package main

import "fmt"

func Strlen(str string) int {
	count := 0
	for _, ch := range str {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			count++
		}
		continue
	}
	return count
}

func main() {
	fmt.Println(Strlen("Alice nn Anne"))
	fmt.Println(Strlen("how are 1 you"))
}