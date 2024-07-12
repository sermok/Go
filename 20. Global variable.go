// Program to illustrate the global variables
package main

import "fmt"

// the global variable declaration
var myvariable1 int = 120

func main() { // from here the local level scope starts
	// the local variables inside the main function
	var myvariable2 int = 210
	// Display value of global variable
	fmt.Printf("Value of Global myvariable1 is : %d\n", myvariable1)
	// Display value of local variable
	fmt.Printf("Value of Local myvariable2 is : %d\n", myvariable2)
	// calling function
	display()
} // local level scope ends
// taking function
func display() { // the local level starts
	// Display value of global variable
	fmt.Printf("Value of Global myvariable1 is : %d\n", myvariable1)
} // the local scope ends here
