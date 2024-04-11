// Program to read and write files
package main

// importing requires packages
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filename string, text string) {
	// fmt package implements formatted I/O
	// and contains the inbuilt methods like the Printf and Scanf
	fmt.Printf("Writing to a file in the Go lang\n")
	// Creating file using Create() method with user inputted filename and err
	// variable cathes any error thrown
	file, er := os.Create(filename)
	if er != nil {
		log.Fatalf("failed creating file: %s", er)
	}
	// closing running file after the main method has completed execution and
	// writing to the file is complete
	defer file.Close()
	// writing data to file using
	// WriteString() method and
	// length of the string is stored in the len variable
	len, er := file.WriteString(text)
	if er != nil {
		log.Fatalf("failed writing to file: %s", er)
	}
	fmt.Printf("\nFile Name is: %s", file.Name())
	fmt.Printf("\nLength is: %d bytes", len)
}
func ReadFile(filename string) {
	fmt.Printf("\n\nReading a file in the Go lang\n")
	// file is read using ReadFile() method of the ioutil package
	data, er := ioutil.ReadFile(filename)
	// in case of an error
	// the error statement is printed, program is stopped
	if er != nil {
		log.Panicf("failed reading data from file: %s", er)
	}
	fmt.Printf("\nFile Name is: %s", filename)
	fmt.Printf("\nSize is: %d bytes", len(data))
	fmt.Printf("\nData is: %s", data)
}

// main function
func main() {
	// user input for the filename
	fmt.Println("Enter-filename: ")
	var filename string
	fmt.Scanln(&filename)
	// user input for the file content
	fmt.Println("Enter-text: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	// file is created then read
	CreateFile(filename, input)
	ReadFile(filename)
}
