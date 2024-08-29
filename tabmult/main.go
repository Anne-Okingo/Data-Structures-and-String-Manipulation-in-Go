// Write a program that displays a number's multiplication table.
// The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.
package main

import (
"os"
"github.com/01-edu/z01"
"strconv"
"fmt"
)

// func main(){
// 	args := os.Args

// 	if len(args) != 2 {
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	num, _ := strconv.Atoi(args[1])

// 	for i := 1; i <= 9; i++{
// 		fmt.Printf("%v * %v = %v\n",i ,num, i*num)
// 	}
// }

func main(){
	args := os.Args

	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}
	

	num, _ := strconv.Atoi(args[1])

	for i := 1; i <= num; i++{
		res := i*num
		result := strconv.Itoa(i) + "*" + strconv.Itoa(num) + "=" + strconv.Itoa(res) + "\n" + " "
		fmt.Println(result)
	}
}