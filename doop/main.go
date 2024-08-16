// Instructions

// Write a program that is called doop.

// The program has to be used with three arguments:

//     A value
//     An operator, one of : +, -, /, *, %
//     Another value

// In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

// The program has to handle the modulo and division operations by 0 as shown on the output examples below.
// Usage

// $ go run .
// $ go run . 1 + 1 | cat -e
// 2
// $
// $ go run . hello + 1
// $ go run . 1 p 1
// $ go run . 1 / 0 | cat -e
// No division by 0
// $
// $ go run . 1 % 0 | cat -e
// No modulo by 0
// $
// $ go run . 9223372036854775807 + 1
// $ go run . -9223372036854775809 - 3
// $ go run . 9223372036854775807 "*" 3
// $ go run . 1 "*" 1
// 1
// $ go run . 1 "*" -1
// -1
// $

// package main

// import (
// 	"fmt"
// 	"os"
// 	// "fmt"
// )

// func Atoi2(str string) int {
// 	result := 0
// 	multiplier := 1

// 	if str[0] == '-' {
// 		multiplier = -1
// 		str = str[1:]
// 	} else if str[0] == '+' {
// 		str = str[1:]
// 	}
// 	for _, ch := range str {
// 		if ch <= '0' || ch >= '9' {
// 			os.Exit(0) // return
// 		}
// 		result = result*10 + int(ch-'0')
// 	}
// 	return result * multiplier
// }

// func Itoa(nb int) string {
// 	sign := ""
// 	if nb < 0 {
// 		sign = "-"
// 		nb = -nb
// 	}
// 	result := ""
// 	for nb != 0 {
// 		mod := nb % 10
// 		startrune := '0'
// 		for i := 0; i < mod; i++ {
// 			startrune++
// 		}
// 		result = string(startrune) + result
// 		nb = nb / 10
// 	}
// 	// fmt.Println(result)
// 	return sign + result
// }

// func main() {
// 	args := os.Args

// 	if len(args) != 4 {
// 		return
// 	}

// 	value1 := Atoi2(args[1])
// 	operator := args[2]
// 	value2 := Atoi2(args[3])
// 	result := 0

	// if operator == "%" && value2 == 0 {
	// 	os.Stdout.WriteString("No Modulo 0\n")
	// 	return
	// } else if operator == "/" && value2 == 0 {
	// 	os.Stdout.WriteString("No division 0\n")
	// 	return
	// }

// 	if value1 >= 9223372036854775807 || value1 <= -9223372036854775807 {
// 		os.Exit(0)
// 		return
// 	} else if value2 >= 9223372036854775807 || value2 <= -9223372036854775807 {
// 		os.Exit(0)
// 	} else if operator == "+" {
// 		result = value1 + value2
// 	} else if operator == "-" {
// 		result = value1 - value2
// 	} else if operator == "*" {
// 		result = value1 * value2
// 		// fmt.Println(result)
// 	} else if operator == "/" {
// 		result = value1 / value2
// 	} else if operator == "%" {
// 		result = value1 % value2
// 	} else {
// 		return
// 	}

// 	// for _, ch := range value1{
// 	// 	if ch < '0' || ch > '9'{
// 	// 		return
// 	// 	}
// 	//}

// 	// results := Itoa(result)
// 	// if result <= -9223372036854775807 || result >= 9223372036854775807 {
// 	// 	return
// 	// } else {
// 	// os.Stdout.WriteString(Itoa(result))
// 	fmt.Print(Itoa((result)))
// 	// }
// 	os.Stdout.WriteString("\n")
// }

package main

import (
	// "fmt"
	"os"
)

func main() {
	args := os.Args

	value1 := Atoi(args[1])
	operator := args[2]
	value2 := Atoi(args[3])
	result := 0

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "%" {
		return
	}
	if operator == "%" && value2 == 0{
		os.Stdout.WriteString("No Modulo of 0")
		return
	}
	if operator == "/" && value2 == 0{
		os.Stdout.WriteString("No division of 0")
		return
	}
	if value1 >= 9223372036854775807 || value1 <= -9223372036854775807{
		os.Exit(0)
	}
	if value2 >= 9223372036854775807 || value2 <= -9223372036854775807{
		os.Exit(0)
	}

	if operator == "+"{
		result = value1 + value2
	}else if operator == "-"{
		result = value1 - value2
	}else if operator == "/"{
		result = value1 / value2
	}else if operator == "*"{
		result = value1 * value2
	}else if operator == "%"{
		result = value1 % value2
	}

	os.Stdout.WriteString(Itoa(result))

	


}

func Atoi(str string)int{
	multiplier := 1
	result := 0

	for _, c := range str{
		if c == '-' && string(str[0]) == "-"{
			multiplier = -1
			str = str[1:]
			continue
		}else if c == '+' && string(str[0]) == "+"{
			multiplier = 1
		}
		if c >= 'a' && c <= 'z'{
			os.Exit(0)
		}
		result = result * 10 + int(c - '0')
	}
	return multiplier * result
}


func Itoa(nb int) string {
	s := ""
	result := ""

	if nb == 0 {
		return "0"
	}

	if nb < 0 {
		s = "-"
		nb = -nb
	}

	for nb > 0 {
		mod := nb % 10
		start := '0'
		for i := 0; i < mod; i++ {
			start++
		}
		result = string(start) + result
		nb = nb / 10
	}
	return s + result
}
