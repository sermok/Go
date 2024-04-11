// Program to illustrate how to rename,
// move a file in the default directory
package main

import (
	"log"
	"os"
)

func main() {
	// Rename and Remove a file
	// Using Rename() function
	OriginalPath := "helloo.txt"
	NewPath := "abc.txt"
	es := os.Rename(Original_Path, New_Path)
	if es != nil {
		log.Fatalf(es)
	}
}
