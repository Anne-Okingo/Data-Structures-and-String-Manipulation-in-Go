// fromto
// Instructions

// Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

//     The numbers must be separated by a comma and a space.
//     If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
//     Prepend a 0 to any number that is less than 10.
//     Add a new line \n at the end of the string.

// Expected function

// func FromTo(from int, to int) string {

// }

// Usage

// Here is a possible program to test your function:

// package main

// import (
// 	"fmt"
// 	"psicine"
// )

// func main() {
// 	fmt.Print(piscine.FromTo(1, 10))
// 	fmt.Print(piscine.FromTo(10, 1))
// 	fmt.Print(piscine.FromTo(10, 10))
// 	fmt.Print(piscine.FromTo(100, 10))
// }

// And its output:

// $ go run . | cat -e
// 01, 02, 03, 04, 05, 06, 07, 08, 09, 10$
// 10, 09, 08, 07, 06, 05, 04, 03, 02, 01$
// 10$
// Invalid$

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(FromTo(1, 10))
// 	fmt.Print(FromTo(10, 1))
// 	fmt.Print(FromTo(10, 10))
// 	fmt.Print(FromTo(100, 10))
	
// }

// func FromTo(from int, to int) string {
// if from > 99 || from < 0 || to > 99 || to < 0 {
// 	return "Invalid\n"
// }
// result := ""
// if from < to {
// 	for i := from; i <= to; i++{
// 		if i > from{
// 			result += ", "
// 		}
// 		if i < 10 {
// 			result =result + "0" + Itoa(i)
// 		}else{
// 			result = result + Itoa(i)
// 		}
// 	}
// 	}else if from > to{
// 		for i := from; i >= to; i--{
// 			if i < from{
// 				result += ", "
// 			}
// 			if i < 10{
// 				result = result + "0" + Itoa(i)
// 			}else{
// 				result+= Itoa(i)
// 			}
// 		}
// 	}else {
// 		result = result + Itoa(from)
// 	}
// 	return result + "\n"
// }

// func Itoa(n int)string{
// 	sign := ""
// 	if n < 0{
// 		sign = "-"
// 		n = -n
// 	}
// 	if n == 0{
// 		return "0"
// 	}

// 	result := ""
// 	for n > 0{
// 		mod := n % 10
// 		startrune := '0'
// 		for i := 0; i < mod; i++{
// 			startrune++
// 		}
// 		result = string(startrune) + result
// 		n= n/10
// 	}
// 	return sign + result
// }

package main

import (
	"fmt"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {
if from > 99 || from < 0 || to > 99 || to < 0{
	return "Invalid\n"
}
result := ""
if from < to{
	for i := from; i <= to ; i++{
		if i > from {
			result = result + ", "
		}
		if i < 10{
			result = result + "0" + Itoa(i)
		}
		if i >= 10 {
			result = result + Itoa(i)
		}
	}
}
if from == to {
	return Itoa(from) + "\n"
}
if  from > to{
	for i := from ; i >= to; i--{
		if i < from{
			result = result + ", "
		}
		if i < 10 {
			result = result + "0" + Itoa(i)
		}
		if i >= 10{
			result = result + Itoa(i)
		}
	}
}
return result + "\n"
}

func Itoa(nb int)string{
	sign := ""
	if nb < 0{
		sign = "-"
		nb = -nb
	}
	if nb == 0{
		return "0"
	}
	result := ""
	for nb > 0{
		mod := nb % 10
		start := '0'
		for i := 0; i < mod; i++{
			start++
		}
		result = string(start) + result
		nb= nb/10
	}
	return sign + result
}
