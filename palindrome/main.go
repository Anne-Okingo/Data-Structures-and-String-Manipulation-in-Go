package main

import (
	"fmt"
)

func main() {
	word := "Ha%nn  aH1"
	Clean := Clean(word)
	fmt.Println(IsPalindrome(Clean))
}

func Clean(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			result += string(ch)
		}
	}
	result2 := ""
	for _, c := range result {
		if c >= 'a' && c <= 'z' {
			result2 += string(c - 32)
		} else {
			result2 += string(c)
		}
	}
	fmt.Println(result2)
	return result2
}

func IsPalindrome(s string) bool {
	s2 := ""
	for i := len(s) - 1; i >= 0; i-- {
		s2 += string(s[i])
	}

	// 	i, j := 0,0

	//	for i < len(s)-1 && j < len(s2)-1{
	//		if s[i] == s2[j]{
	//			i++
	//		}
	//		j++
	//	}
	//	if i == len(s)-1 && j == len(s2)-1{
	//		return true
	//	}
	//
	// return false
	return s == s2
}
