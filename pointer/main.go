package main

import (
	"fmt"
)

func main(){
	a, b := 42, 2701

	p := &a
	fmt.Println(*p)

	*p = 21
	fmt.Println(*p)

	*p = *p / 2
	fmt.Println(*p)

	l := &b
	fmt.Println(*l)
}


// package main

// import "fmt"

// func main() {
//     i, j := 42, 2701

//     // Declaring a pointer variable p which holds the memory address of an integer
//     p := &i         // p now points to the memory address of i
//     fmt.Println(*p) // Print the value pointed to by p, which is the value of i
//     *p = 21         // Set the value of i through the pointer p
//     fmt.Println(i)  // Print the new value of i

//     p = &j          // Now p points to the memory address of j
//     *p = *p / 37    // Divide the value pointed to by p (which is the value of j) by 37
//     fmt.Println(j)  // Print the new value of j
// }


// i, j := 42, 2701: This declares two integer variables i and j and assigns them values 42 and 2701 respectively.

// p := &i: Here, &i takes the memory address of i (the address where i is stored in memory), and assigns it to the pointer variable p. So, p now holds the memory address of i. p is a pointer to i.

// fmt.Println(*p): This line prints the value pointed to by p. The * operator before p dereferences the pointer p, meaning it retrieves the value stored at the memory address pointed to by p, which is the value of i.

// *p = 21: Here, *p is used to set the value at the memory address pointed to by p (which is the value of i) to 21.

// p = &j: Now, p is reassigned to point to the memory address of j.

// *p = *p / 37: This line divides the value pointed to by p (which is the value of j) by 37 and stores the result back in the memory location pointed to by p (which is j).

// fmt.Println(j): Finally, this line prints