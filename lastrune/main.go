// package main

// import "fmt"

// func main(){
// 	fmt.Println(LastRune("Alice"))
// }

// func LastRune(str string)string{
// 	var news string
// 	for range str{
// 		news = string(str[len(str)-1])
// 	}
// 	return news
// }

package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(LastRune("Alice Anne"))
	z01.PrintRune('\n')
}

func LastRune(str string)rune{
	s := []rune(str)
	return s[len(s)-1]
}

