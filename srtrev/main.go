package main

import "fmt"

func main(){
	fmt.Println(StrRev("Alice Anne"))
}

func StrRev(str string)string{
	result := ""
	for _, ch := range str{
		result = string(ch) + result
	}
	return result
}