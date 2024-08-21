package main

import "fmt"

func StrReplace(str, str1, str2 string) string{
	new := ""
	for _,ch := range str{
		for _, chr := range str1 {
			if ch == chr {
				new = new + str2
			}else {
				new = new + string(ch)
			}
		}
	}
	return new
}

func main() {
	fmt.Println(StrReplace("hey anne", "a", "A"))
}