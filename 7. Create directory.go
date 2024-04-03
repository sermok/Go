package main

import (
	"log"
	"os"
)

func main() {
	if er := os.Mkdir("a", os.ModePerm); er != nil {
		log.Fatal(er)
	}
}
