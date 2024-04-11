package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// os.Open() opens specific file in the
	// read-only mode,
	// this return pointer of type os.
	file, er := os.Open("sample1.txt")
	if er != nil {
		log.Fatalf("failed to open")
	}
	// bufio.NewScanner() function is called in which
	// object os.File passed as its parameter
	// this returns object bufio.Scanner which is used on the
	// bufio.Scanner.Split() method
	scanner := bufio.NewScanner(file)
	// The bufio.ScanLines is used as
	// input to method bufio.Scanner.Split()
	// and then scanning forwards to each
	// new line using bufio.Scanner.Scan() method.
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	// The method os.File.Close() is called
	// on the os.File object to close file
	file.Close()
	// and then a loop iterate through,
	// prints each of the slice values.
	for _, each_ln := range text {
		fmt.Println(each_ln)
	}
}
