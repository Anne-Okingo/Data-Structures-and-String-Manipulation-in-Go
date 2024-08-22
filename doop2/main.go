package main

import (
	// "fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		os.Exit(0)
	}
	value1 := Atoi(args[1])
	operator := args[2]
	value2 := Atoi(args[3])
	result := 0

	if operator != "*" && operator != "-" && operator != "+" && operator != "/" && operator != "%" {
		os.Exit(0)
	}
	if value1 >= 9223372036854775807 || value1 <= -9223372036854775807{
		os.Exit(0)
	}
	if value2 >= 9223372036854775807 || value2 <= -9223372036854775807{
		os.Exit(0)
	}

	if operator == "%" && value2 == 0{
		os.Stdout.WriteString("No modulo by 0")
		return
	}
	if operator == "/" && value2 == 0{
		os.Stdout.WriteString("No division by 0")
		return
	}

	if operator == "+"{
		result = value1 + value2
	}else if operator == "-"{
		result = value1 - value2
	} else if operator == "*"{
		result = value1 * value2
	} else if operator == "/"{
		result = value1 / value2
	} else if operator == "%"{
		result = value1 % value2
	} 
	line := "\n"
	os.Stdout.WriteString(Itoa(result))
	os.Stdout.WriteString(line)

}

func Atoi(str string) int {
	result := 0
	multiplier := 1
	for _, c := range str {
		if c == '-' && string(str[0]) == "-" {
			multiplier = -1
			str = str[1:]
			continue
		} else if c == '+' && string(str[0]) == "+" {
			multiplier = 1
		}
		if c >= 'a' && c <= 'z' {
			os.Exit(0)
		}
		result = result*10 + int(c-'0')
	}
	return multiplier * result
}

func Itoa(nb int) string {
	sign := ""
	result := ""
	if nb < 0 {
		sign = "-"
		nb = -nb
	}
	if nb == 0 {
		return "0"
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
	return sign + result
}
