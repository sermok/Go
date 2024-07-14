// Program to show the compiler preference
// to local variable over a global variable
package main

import "fmt"

// the global variable declaration
var myvariable1 int = 120

func main() { // from here the local level scope starts
	// local variables inside main function
	// it is same as a global variable
	var myvariable1 int = 210
	// Display value
	fmt.Printf("Value of myvariable1 is : %d\n", myvariable1)
} // here the local level scope ends
