// Write a program that prints the arguments received in the command line.
package main

import(
	"os"
	"github.com/01-edu/z01"
)

func main(){
	args := os.Args[1:]
	for _, ch := range (args){
		for _, chr := range ch{
		z01.PrintRune(chr)
		}
		z01.PrintRune('\n')
	}
}