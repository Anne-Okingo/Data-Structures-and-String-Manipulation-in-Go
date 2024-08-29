package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	args := os.Args
	name := args[0]
	// fmt.Println(name)
	for _, hc := range(name){
			z01.PrintRune(hc)
	}
	z01.PrintRune('\n')
}