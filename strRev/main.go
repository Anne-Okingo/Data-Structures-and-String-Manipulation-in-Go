package main

import "fmt"

func StrRev(str string) string {
	rev := []rune{}
	p := []rune(str)
	length := len(p)
	for i := length - 1; i >= 0; i-- {
		rev = append(rev, p[i])
	}
	return string(rev)
}

func main() {
	fmt.Println(StrRev("Alice"))
}
