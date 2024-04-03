package main

import (
	"log"
	"os"
)

var (
	myfile *os.FileInfo
	es     error
)

func main() {
	// Stat() function returns the file info and
	// if there is no file, then it will return error
	myfile, es := os.Stat("/Users/Sergey/Documents/Go/helloo.txt")
	if es != nil {
		// Checking if given file exists or not
		// Using IsNotExist() function
		if os.IsNotExist(es) {
			log.Fatal("File not Found")
		}
	}
	log.Println("File Exist")
	log.Println("File Detail is:")
	log.Println("Name is: ", myfile.Name())
	log.Println("Size is: ", myfile.Size())
}
