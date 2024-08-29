package main

import (
"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(str string)string{
	// if len(str) == 0 || len(str) > 3{
	// 	return "G" + "\n"
	// }else {
	// 	return "Invalid Input"
	// }
	if len(str) == 0{
		return "G" + "\n"
	}
	count := 0

	for range str{
		count++
	}
	if count > 3{
		return "G" + "\n"
	}else{
		return "Invalid Input"
	}
}
