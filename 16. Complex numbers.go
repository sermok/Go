// Illustrate the use of complex numbers
package main

import "fmt"

func main() {
	var x complex128 = complex(7, 3)
	var y complex64 = complex(complex)
	fmt.Println(x)
	fmt.Println(y)
	// Display type
	fmt.Printf("The type of x is %T and "+"the type of y is %T", x, y)
}
