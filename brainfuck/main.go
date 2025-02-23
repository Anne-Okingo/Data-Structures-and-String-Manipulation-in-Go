// brain_fuck
// Instructions
// Write a Brainfuck interpreter program. The source code will be given as first parameter. The code will always be valid, with less than 4096 operations. Brainfuck is a minimalist language. It consists of an array of bytes (in this exercise 2048 bytes) all initialized with zero, and with a pointer to its first byte.

// Every operator consists of a single character :

// '>' increment the pointer
// '<' decrement the pointer
// '+' increment the pointed byte
// '-' decrement the pointed byte
// '.' print the pointed byte on standard output
// '[' go to the matching ']' if the pointed byte is 0 (loop start)
// ']' go to the matching '[' if the pointed byte is not 0 (loop end)
// Any other character is a comment.

// For receiving arguments from the command line you should use something like:

// fn main() {
//     let args: Vec<String> = std::env::args().collect();

//     brain_fuck(&args[1]);
// }

// Usage
// $ cargo run "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
// Hello World!
// $ cargo run "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>."
// Hi
// $ cargo run "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++."
// abc


package main

import(
"os"
"github.com/01-edu/z01"
)

func main(){
	args := os.Args

	if len(args) != 2{
		return
	}

	arg := args[1]

	memmorytape := make([]byte,2048)
	pointer := 0
	brackets := make(map[int]int)
	stack := []int{}

	for i, ch := range arg{
		if ch == '['{
			stack = append(stack, i)
		}
		if ch == ']'{
			oppening := stack[len(stack)-1]
			brackets[oppening] = i
			brackets[i] = oppening
			stack = stack[:len(stack)-1]
		}
	}

	for i := 0; i < len(arg); i++{

		if arg[i] == '>'{
			pointer++
		}
		if arg[i] == '<'{
			pointer--
		}
		if arg[i] == '+'{
			memmorytape[pointer]++
		}
		if arg[i] == '-'{
			memmorytape[pointer]--
		}

		if arg[i] == '.'{
			z01.PrintRune(rune(memmorytape[pointer]))
		}

		if arg[i] == '[' && memmorytape[pointer] == 0{
			i = brackets[i]
		}
		if arg[i] == ']' && memmorytape[pointer] != 0{
			i = brackets[i]
		}
	}

}