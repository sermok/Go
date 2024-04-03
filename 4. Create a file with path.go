package main

import (
	"log"
	"os"
)

func main() {
	// empty file creation
	// Create() function Using
	myfile, es := os.Create("/Users/Sergey/Documents/Go/helloo.txt")
	if es != nil {
		log.Fatal(es)
	}
	log.Println(myfile)
	myfile.Close()
}
