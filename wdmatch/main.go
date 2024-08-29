// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

// If the number of arguments is different from 2, the program displays nothing.

// $ go run . 123 123
// 123
// $ go run . faya fgvvfdxcacpolhyghbreda
// faya
// $ go run . faya fgvvfdxcacpolhyghbred
// $ go run . error rrerrrfiiljdfxjyuifrrvcoojh
// $ go run . "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
// quarante deux
// $ go run .
// $

package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// args := os.Args

	if len(os.Args) < 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	if len(s1) > len(s2) {
		return
	}
	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	if i == len(s1) {
			for _, ch := range s1{
				z01.PrintRune(ch)
			}
			z01.PrintRune('\n')
	}
}

