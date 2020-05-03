package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MateoM24/HTML-links-extractor.git/search"
)

func main() {
	const fileName = "testresources/ex2.html"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Cannot open file:", fileName)
	}
	results := search.RetrieveLinks(file)
	fmt.Println(*results)
}
