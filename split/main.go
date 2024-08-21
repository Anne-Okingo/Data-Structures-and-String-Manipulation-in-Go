// Write a function that receives a string and a separator and returns a slice of strings that results of splitting the string s by the separator sep.
package main

import (
	"fmt"
	// "piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	// fmt.Println(len(s))
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	result := []string{}
	seplen := len(sep)
	// to truck the beginning of the index of the current substring
	start := 0

 
	// The loop checks all possible positions in s where the separator sep could be located. ie 21-2=17
	for i := 0; i <= len(s)- seplen; i++{
		// say i is at 0 this will iterate over untill the string at strings at index i but not including i+seplen (s [i:i+seplen]) matchs sep
		if s[i : i + seplen] == sep{
			result = append(result, s[start:i])
			start = i + seplen
		}
	}
	if start <= len(s){
		result = append(result, s[start:])
	}

return result
}



