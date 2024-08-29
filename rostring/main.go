// rostring
// Instructions

// Write a program that takes a string and displays this string after rotating it one word to the left.

// Thus, the first word becomes the last, and others stay in the same order.

// A word is a sequence of alphanumerical characters.

// Words will be separated by only one space in the output.

// If the number of arguments is different from 1, the program displays a newline.
// Usage

// $ go run . "abc   " | cat -e
// abc$
// $ go run . "Let there     be light"
// there be light Let
// $ go run . "     AkjhZ zLKIJz , 23y"
// zLKIJz , 23y AkjhZ
// $ go run . | cat -e
// $
// $

package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) != 2{
		return
	}
	str := os.Args[1]
	Split(str)
}

func Split(s string){
	result := ""
	sliced := []string{}
	for _, ch := range s{
		if ch != ' '{
			result += string(ch)
		}else if result != ""{
			sliced = append(sliced,result)
			result = ""
		}
	}
	if result != ""{
		sliced = append(sliced,result)
	}
	result1 := ""
	
	for i := 1; i < len(sliced); i++{
		result1 += string(sliced[i]) + " "
	}
	result1 += string(sliced[0])
	
	for _, ch := range result1{
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}



