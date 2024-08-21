// swapbits
// Instructions

// Write a function that takes a byte, swaps its halfs (like the example) and returns the result.
// Expected function

// func SwapBits(octet byte) byte {

// }

// Example:

// 1 byte

// 0100 | 0001
//     \ /
//     / \
// 0001 | 0100

package main

import (
	//"fmt"

	"fmt"

	//"github.com/01-edu/z01"
)

func main() {
    var octet byte = 0xAB // 10101011 in binary
    lowerNibble := octet & 0x0F

    fmt.Printf("Original byte: %08b (%d in decimal)\n", octet, octet)
    fmt.Printf("Lower nibble (0x0F): %08b (%d in decimal)\n", lowerNibble, lowerNibble)
}
// a := byte(0b01000100)
// //b := byte(0b00010100)
// swapped := SwapBits(a)
// fmt.Println(swapped)

func SwapBits(octet byte) byte {
a := octet >> 4
b:= octet<< 4
return a | b
}
// func main() {
//     octet := byte(0b01000001)
//     swapped := swapbits(octet)
//     // z01.PrintRune(rune(swapped))
//     // z01.PrintRune('\n')
//     fmt.Println(swapped)
// }

// func swapbits(octet byte)byte{
//     a := octet >> 4 | octet << 4
//     return a
// }