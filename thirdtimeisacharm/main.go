package main

import "fmt"

func ThirdTimeIsACharm(str string)string{
	 if str == "" || len(str) < 3 {
		return ""
	 }

	 result := ""
	 for i := 2; i < len(str); i=i+3{
		result = result + string(str[i])
	 }
	 return result
}

func main(){
	fmt.Println(ThirdTimeIsACharm("An*e"))
	fmt.Println(ThirdTimeIsACharm("An*eiou"))
}