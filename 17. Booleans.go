// Program to illustrate the use of booleans
package main

import "fmt"

func main() {
	// variables
	strg1 := "PeeksofPeeks"
	strg2 := "peeksofpeeks"
	strg3 := "PeeksofPeeks"
	results1 := strg1 == strg2
	results2 := strg1 == strg3
	// Display result
	fmt.Println(results1)
	fmt.Println(results2)
	// Display type of
	// results1 and results2
	fmt.Printf("The type of results1 is %T and "+"the type of results2 is %T", results1, results2)
}
