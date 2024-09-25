// reversestrcap
// Instructions

// Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').

// If there are no argument, the program displays nothing.
// Usage

// $ go run . "First SMALL TesT" | cat -e
// firsT smalL tesT$
// $ go run . "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
// seconD tesT iS A littlE easieR$
// bewarE it'S noT harD wheN $
//  gO A dernieR 0123456789 foR thE roaD E$
// $ go run .
// $

package main

import (
	"fmt"
	"os"
	// "github.com/01-edu/z01"
)
func main(){
	if len(os.Args) < 2{
		return
	}
	slice := os.Args[1:]
	for _, str := range slice{
		result := ReverseStrCap(str)
		fmt.Println(result)
	}
}

func ReverseStrCap(s string)string{
	result := ""
	for i, ch := range s{ 
		if i < len(s)- 1{
			if ToUpper(ch) && s[i +1]!= ' '{
				result += string(ch + 32)
			}else if Tolower(ch) && s[i+1] == ' '{
				result += string(ch-32)
			}else{
				result += string(ch)
			}
		}else if Tolower(ch) && i == len(s)-1{
			result += string(ch-32)
		}else{
			result += string(ch)
		}
		}
		return result
}

func Tolower(r rune)bool{
	return r >= 'a' && r <= 'z'
}

func ToUpper(r rune)bool{
	return r >= 'A' && r <= 'Z'
}