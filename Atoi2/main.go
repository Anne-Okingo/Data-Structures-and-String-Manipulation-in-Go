// package main

// import "fmt"

// func Atoi2(str string)int{
// result := 0
// multiplier := 1
// for i, ch := range str{
// 	if ch == '-' && i == 0 {
// 		multiplier = -1
// 		continue
// 	}else if ch == '+' && i == 0 {
// 		multiplier = 1
// 		continue
// 	}

// 	result = (result * 10 + int(ch - '0'))
// }
// //fmt.Println(result)
// return result * multiplier
// }


// func main(){
// 	fmt.Println(Atoi2("-254"))
// }

package main

import(
"fmt"
)
 func main(){
	fmt.Println(Atoi2("2546"))
 }

 func Atoi2(str string)int{
	result := 0
	multiplier := 1

	for _, ch := range str{
		if ch == '-' && string(str[0]) == "-"{
			multiplier = -1
			continue
		}else if ch == '+' && string(str[0]) == "+"{
			multiplier = 1
		}
		result = result * 10 + int(ch - '0')
	}
	return multiplier * result
 }