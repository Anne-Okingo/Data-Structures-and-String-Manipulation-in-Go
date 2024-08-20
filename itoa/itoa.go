package main

import "github.com/01-edu/z01"
// "fmt"

func main() {
	result := ""
	num := 123

	for num != 0 {
		mod := num % 10
		firstrune := '0'
		for i := 0; i < mod; i++ {
			firstrune++
		}
		result = result + string(firstrune)
		num = num / 10
	}
	for j := len(result)- 1; j >= 0; j--{
		z01.PrintRune(rune(result[j]))
		// fmt.Println(string(result[j]))
	}
}


// package main

// import "fmt"

// func Itoa(nb int)string{
// 	result := ""
// 	for nb != 0{
// 		mod := nb % 10
// 		firstrune := '0'
// 		for i := 0; i < mod; i++ {
// 			firstrune++
// 		}
// 		result = result + string(firstrune)
// 		nb = nb/10
// 	}
// 	reversed := ""
// 	for i := len(result)-1; i >= 0; i--{
// 		reversed = reversed + string(result[i])
// 	}
// 	return reversed
// }

// func main(){
// 	fmt.Println(Itoa(254))
// }