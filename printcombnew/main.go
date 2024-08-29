// printcomb
// Instructions

// Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.

// These combinations are separated by a comma and a space.
// Expected function

// func PrintComb() {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import "piscine"

// func main() {
// 	piscine.PrintComb()
// }

// This is the incomplete output :

// $ go run . | cat -e
// 012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789$
// $

//     000 or 999 are not valid combinations because the digits are not different.

//     987 should not be shown because the first digit is not less than the second.

// Notions

//     01-edu/z01

package main

import(
"github.com/01-edu/z01"
)

func main(){
for i := '0';i <= '7';i++{
	for j := '1'; j <= '8'; j++{
		for k := '2'; k <= '9'; k++{
			if i > j || j > k || i== j || j == k {
				continue
			}else if i == '7' && j == '8'&& k == '9'{
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				z01.PrintRune('\n')
			}else{
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
}

