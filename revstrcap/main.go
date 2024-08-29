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

import(
"os"
"fmt"
)

func main(){
	args := os.Args[1:]
	 if len(args) < 1{
		return
	 }

	 for _, str := range args{
		words := ""
		for _, ch := range str{
			if ch >= 'A' && ch <= 'Z' {
				words += string(ch + 32) 
			}else{
				words += string(ch)
			}
		}
		// fmt.Println(words)
		result := ""
		for i := 0; i < len(words); i++{
			if i == len(words)-1  && words[i] >= 'a' && words[i] <= 'z' || i < len(words)-1 && words[i+1] == ' ' && words[i] >= 'a' && words[i] <= 'z' {
				result += string(words[i]-32)
			}else{
				result += string(words[i])
			}
		}
		fmt.Println(result)
		 }
}