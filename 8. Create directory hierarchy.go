package main

import (
	"log"
	"os"
)

func main() {
	if er := os.MkdirAll("a/b/c/d", os.ModePerm); er != nil {
		log.Fatal(er)
	}
}
