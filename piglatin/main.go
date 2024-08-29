package main 

 import (
"os"
"fmt"
 )

 func main(){
	args := os.Args
	if len(args) != 1{
		return
	}
	word := args[1]

	piglatined := Piglatin(word)

	if piglatined == ""{
		fmt.Println("No Vowels")
	}else {
		fmt.Println(piglatined)
	}
 }

 func Piglatin(str string)string{
	for i, ch := range str{
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			if i == 0{
				return str + "ay"
			}
				return str[i:] + str[:i] + "ay"
		}
	}
	return ""
 }