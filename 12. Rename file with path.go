// Program to illustrate how to rename,
// remove a file in new directory
package main

import (
	"log"
	"os"
)

func main() {
	// Rename and Remove file
	// Using Rename() function
	OriginalPath := "/Users/Sergey/Documents/Go/helloo.txt"
	NewPath := "/Users/Sergey/Documents/Go/abc.txt"
	es := os.Rename(OriginalPath, NewPath)
	if es != nil {
		log.Fatalf(es)
	}
}
