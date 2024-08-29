revwstr
Instructions

Write a program that takes a string as a parameter, and prints its words in reverse, followed by a newline.

    A word is a sequence of alphanumerical characters.

    If the number of arguments is different from 1, the program will display nothing.

    In the parameters that are going to be tested, there will not be any extra spaces. (meaning that there will not be additional spaces at the beginning or at the end of the string and that words will always be separated by exactly one space).

Usage

$ go run . "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
$ go run . "abcdefghijklm"
abcdefghijklm
$ go run . "he stared at the mountain"
mountain the at stared he
$ go run . "" | cat-e
$
$




hashcode
Instructions

Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

    The hash equation is computed as follows:

(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

    If the resulting character is unprintable add 33 to it.

Expected function

func HashCode(dec string) string {
}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

And its output:

$ go run .
B
CD
EDF
Spwwz+bz}wo





printmemory
Instructions

Write a function that takes (arr [10]byte), and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.
Expected function

func PrintMemory(arr [10]byte) {

}

Usage

Here is a possible program to test your function :

package main

import "piscine"

func main() {
	piscine.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

And its output :

$ go run . | cat -e
68 65 6c 6c$
6f 10 15 2a$
00 00$
hello..*..$
$





cameltosnakecase
Instructions

Write a function that converts a string from camelCase to snake_case.

    If the string is empty, return an empty string.
    If the string is not camelCase, return the string unchanged.
    If the string is camelCase, return the snake_case version of the string.

For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

    lowerCamelCase
    UpperCamelCase

Rules for writing in camelCase:

    The word does not end on a capitalized letter (CamelCasE).
    No two capitalized letters shall follow directly each other (CamelCAse).
    Numbers or punctuation are not allowed in the word anywhere (camelCase1).

Expected function

func CamelToSnakeCase(s string) string{

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

And its output:

$ go run .
Hello_World
hello_World
camel_Case
CAMELtoSnackCASE
camel_To_Snake_Case
hey2






weareunique
Instructions

Write a function that takes two strings's' and returns the number of characters that are not included in both, without repeating characters.

    If there is no unique characters return 0.
    If both strings are empty return -1.

Expected function

func WeAreUnique(str1 , str2 string) int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

And its output:

$ go run .
2
-1
6



Can Jump

Given an array of non-negative integers representing the number of steps you can take forward from each position, implement the function CanJump() which takes a slice of unsigned integers []uint as input and returns a boolean value. This function should determine if it's possible to reach and stay at the last index of the array starting from the first index, based on the steps you need to advance. Be aware that:

    Each value represents the exact number of steps you must take forward from that position.
    The function should return true if it's possible to reach and stay at the last index without stepping out of the array, and false otherwise.
    If the input has only one element, that is the last position in the array so the function will return true but if the array is empty it returns false.

Let's take the example array input := []uint{2, 3, 1, 1, 4}.


Position: 0  1  2  3  4
Steps:    2  3  1  1  4
          ^

// Starting from position 0, you have 2 steps to move forward. This means you will move to positions 2.

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                ^

// From position 2, you have 1 step, so you will move to position 3.

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                   ^

// Finally, from position 3, you have 1 step to reach the last index at position 4 confirming that it's possible so the output will be "True".

Position: 0  1  2  3  4
Steps:    2  3  1  1  4
                      ^

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []uint{0}
	fmt.Println(piscine.CanJump(input3))
}

And its output :

$ go run .
true
false
true
$


fishandchips
Instructions

Write a function called FishAndChips() that takes an int and returns a string.

    If the number is divisible by 2, print fish.
    If the number is divisible by 3, print chips.
    If the number is divisible by 2 and 3, print fish and chips.
    If the number is negative return error: number is negative.
    If the number is non divisible by 2 or 3 return error: non divisible.

Expected function

func FishAndChips(n int) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FishAndChips(4))
	fmt.Println(piscine.FishAndChips(9))
	fmt.Println(piscine.FishAndChips(6))
}

And its output:

$ go run . | cat -e
fish$
chips$
fish and chips$



countalpha
Instructions

Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.
Expected functions

func CountAlpha(s string) int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

And its output:

$ go run .
10
5
5




slice
Instructions

The function receives a slice of strings and one or more integers, and returns a slice of strings. The returned slice is part of the received one but cut from the position indicated in the first int, until the position indicated by the second int.

In case there only exists one int, the resulting slice begins in the position indicated by the int and ends at the end of the received slice.

The integers can be negative.
Expected function

func Slice(a []string, nbrs... int) []string{

}

Usage

Here is a possible program to test your function :

package main

import (
    "fmt"
    "piscine"
)

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", piscine.Slice(a, 1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
    fmt.Printf("%#v\n", piscine.Slice(a, -3))
    fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))
}

$ go run .
[]string{"algorithm", "ascii", "package", "golang"}
[]string{"ascii", "package"}
[]string{"ascii", "package", "golang"}
[]string{"package"}
[]string(nil)




printrevcomb
Instructions

Write a program that prints in descending order on a single line all unique combinations of three different digits so that the first digit is greater than the second and the second is greater than the third.

These combinations are separated by a comma and a space.
Usage

Here is an incomplete output :

$ go run . | cat -e
987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
$

999 or 000 are not valid combinations because the digits are not different.

789 should not be shown because the first digit is not greater than the second.



package main

import (
	"os"
)

func main() {
	// Create a byte slice to build the output
	var result []byte

	// Generate all unique combinations of three different digits
	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				result = append(result, '0'+byte(i))
				result = append(result, '0'+byte(j))
				result = append(result, '0'+byte(k))
				result = append(result, ',')
				result = append(result, ' ')
			}
		}
	}

	// Remove the last ", "
	if len(result) > 1 {
		result = result[:len(result)-2]
	}

	// Write the result to stdout
	os.Stdout.Write(result)
}




retainfirsthalf
Instructions

Write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.

    If the length of the string is odd, round it down.
    If the string is empty, return an empty string.
    If the string length is equal to one, return the string.

Expected function

func RetainFirstHalf(str string) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
}

And its output:

$ go run . | cat -e
This is the 1st half$
A$
$
Hello$



expandstr
Instructions

Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

The string will be followed by a newline ('\n').

A word, in this exercise, is a sequence of visible characters.

If the number of arguments is not 1, or if there are no word, the program displays nothing.
Usage

$ go run . "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
$ go run . "   only  it's harder   " | cat -e
only   it's   harder$
$ go run . " how funny it is" "did you  hear, Mathilde ?" | cat -e
$ go run .
$





zipstring
Instructions

Write a function that takes a string and returns a new string that replaces every character with the number of duplicates and the character itself, deleting the extra duplications.

    The letters are from the latin alphabet list only. Any other character, symbols, shall not be tested.

Expected function

func ZipString(s string) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

And its output:

$ go run .
1Y1o3u1n1g1F1e4l1a1s
1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
1H1e2l2o1 1T1h1e2r1e1!





secondhalf
Instructions

Write a function SecondHalf() that receives a slice of int and returns another slice of int with the second half of the values.

    If the length is odd, include the middle value in the result.

Expected function

func SecondHalf(slice []int) []int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SecondHalf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(piscine.SecondHalf([]int{1, 2, 3}))
}

And its output:

$ go run . | cat -e
[6 7 8 9 10]
[2 3]





fprime
Instructions

Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

    Factors must be displayed in ascending order and separated by *.

    If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

Usage

$ go run . 225225
3*3*5*5*7*11*13
$ go run . 8333325
3*3*5*5*7*11*13*37
$ go run . 9539
9539
$ go run . 804577
804577
$ go run . 42
2*3*7
$ go run . a
$ go run . 0
$ go run . 1
$



concatslice
Instructions

Write a function ConcatSlice() that takes two slices of integers as arguments and returns the concatenation of the two slices.
Expected function

func ConcatSlice(slice1, slice2 []int) []int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{}))
}

And its output:

$ go run .
[1 2 3 4 5 6]
[4 5 6 7 8 9]
[1 2 3]




options
Instructions

Write a program that takes an undefined number of arguments which could be considered as options and writes on the standard output a representation of those options as groups of bytes followed by a newline ('\n').

    An option is an argument that begins with a - and that can have multiple characters which could be : -abcdefghijklmnopqrstuvwxyz

    All options are stocked in a single int and each options represents a bit of that int, and should be stocked like this :

          - 00000000 00000000 00000000 00000000
          - ******zy xwvutsrq ponmlkji hgfedcba

    Launching the program without arguments or with the -h flag activated must print all the valid options on the standard output, as shown on one of the following examples.

    Please note the -h flag has priority over the others flags when it is called first in one of the arguments. (See the examples)

    A wrong option must print Invalid Option followed by a newline.

Usage

$ go run . | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -abc -ijk | cat -e
00000000 00000000 00000111 00000111$
$ go run . -z | cat -e
00000010 00000000 00000000 00000000$
$ go run . -abc -hijk | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -h | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -zh | cat -e
00000010 00000000 00000000 10000000$
$ go run . -z -h | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -hhhhhh | cat -e
options: abcdefghijklmnopqrstuvwxyz$
$ go run . -eeeeee | cat -e
00000000 00000000 00000000 00010000$
$ go run . -% | cat -e
Invalid Option$
$ go run . - | cat -e
Invalid Option$
$




revconcatalternate
Instructions

Write a function RevConcatAlternate() that receives two slices of int as arguments and returns a new slice with alternated values of each slice in reverse order.

    The input slices can have different lengths.
    The new slice should start with the elements from the largest slice first and when they became equal size slices, it should add an element of the first given slice.
    If the slices are of equal length, the new slice should start with an element of the first slice.

    Note: you can check the examples bellow for more details.

Expected function

func RevConcatAlternate(slice1,slice2 []int) []int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

And its output:

$ go run .
[3 6 2 5 1 4]
[9 8 7 3 6 2 5 1 4]
[8 9 3 2 5 1 4]
[3 2 1]






Concat Str
Instructions

Create a function named concatStr which takes 2 arguments and concatenates them.



concatparams
Instructions

Write a function that takes the arguments received in parameters and returns them as a string. The string is the result of all the arguments concatenated with a newline (\n) between.
Expected function

func ConcatParams(args []string) string {

}

Usage

Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
}

And its output :

$ go run .
Hello
how
are
you?
$




saveandmiss
Instructions

Write a function called SaveAndMiss() that takes a string and an int as an argument. The function should move through the string in sets determined by the int, saving the first set, omitting the second, saving the third, and so on, in a 'save' and 'miss' fashion until the end of the string is reached. Return a string containing the saved characters.

    If the int is 0 or a negative number return the original string.

Expected function

func SaveAndMiss(arg string, num int) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SaveAndMiss("123456789", 3))
	fmt.Println(piscine.SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(piscine.SaveAndMiss("", 3))
	fmt.Println(piscine.SaveAndMiss("hello you all ! ", 0))
	fmt.Println(piscine.SaveAndMiss("what is your name?", 0))
	fmt.Println(piscine.SaveAndMiss("go Exercise Save and Miss", -5))
}

And its output:

$ go run . | cat -e
123789$
abcghimnostuz$
$
hello you all ! $
what is your name?$
go Exercise Save and Miss$






	wordflip
Instructions

Write a function WordFlip() that takes a string as input and returns it in reverse order.

    The output should be followed by a newline \n.
    If the string is empty, return Invalid Output.
    Ignore multiple spaces between words and trim any leading or trailing spaces in the string.

Expected function

func WordFlip(str string) string {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.WordFlip("First second last"))
	fmt.Print(piscine.WordFlip(""))
	fmt.Print(piscine.WordFlip("     "))
	fmt.Print(piscine.WordFlip(" hello  all  of  you! "))
}

And its output:

$ go run . | cat -e
last second First$
Invalid Output$
$
you! of all hello$




concatalternate
Instructions

Write a function ConcatAlternate() that receives two slices of an int as arguments and returns a new slice with the result of the alternated values of each slice.

    The input slices can be of different lengths.
    The new slice should start with an element of the largest slice.
    If the slices are of equal length, the new slice should return the elements of the first slice first and then the elements of the second slice.

Expected function

func ConcatAlternate(slice1,slice2 []int) []int {

}

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

And its output:

$ go run .
[1 4 2 5 3 6]
[1 2 3 4 5 6 7 8 9 10 11]
[4 1 5 2 6 3 7 8 9]
[1 2 3]