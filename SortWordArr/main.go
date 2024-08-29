// Write a function SortWordArr that sorts by ascii (in ascending order) a string slice.

package main

import (
	"fmt"
	// "piscine"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(a []string) {
	n := len(a)

	for i := 0; i < n-1; i++{
		for j := 0; j < n-i-1; j++{
			if a[j] > a[j+1]{
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

