// Program to illustrate the local variables
package main

import "fmt"

// main function
func main() { // from here the local level scope of main function starts
	// local variables inside the main function
	var myvariable1, myvariable2 int = 90, 47
	// Display values of the variables
	fmt.Printf("Value of myvariable1 is : %d\n", myvariable1)
	fmt.Printf("Value of myvariable2 is : %d\n", myvariable2)
} // here the local level scope of main function ends
