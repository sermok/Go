// program to read and write files
package main

// importing packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {
	// fmt package implements formatted I/O, it has functions like Printf and Scanf
	fmt.Printf("Writing file in Go lang/n")
	// in case error is thrown it is received by err variable and Fatalf method of
	// log prints error message and stops program execution
	file, er := os.Create("test1.txt")
	if er != nil {
		log.Fatalf("failed creating file: %s", er)
	}
	// Defer is used for the purposes of cleanup like closing a running file after the file has
	// been written and the main function has completed execution
	defer file.Close()
	// len variable cartures the length of string written to the file.
	len, er := file.WriteString("Welcome Everyone" + "Program demonstrates reading and writing" + "operations to a file in the Go lang.")
	if er != nil {
		log.Fatalf("failed writing to file: %s", er)
	}
	// Name() method returns name of the file as presented to Create() method.
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}
func ReadFile() {
	fmt.Printf("\n\nReading a file in the Go lang\n")
	fileName := "test1.txt"
	// The ioutil package contains inbuilt
	// methods like ReadFile that reads
	// filename and returns contents.
	data, er := ioutil.ReadFile("test.txt")
	if er != nil {
		log.Panicf("failed reading data from file: %s", er)
	}
	fmt.Printf("\nFile Name is: %s", fileName)
	fmt.Printf("\nSize is: %d bytes", len(data))
	fmt.Printf("\nData is: %s", data)
}

// main function
func main() {
	CreateFile()
	ReadFile()
}
