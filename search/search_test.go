package search

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestRetrieveLinks(t *testing.T) {
	scenarios := prepareScenarios()
	for k, e := range scenarios {
		results := RetrieveLinks(k)
		if len(*results) != len(e) {
			t.Errorf("Found %d number of records but should have found %d", len(*results), len(e))
		}

	}
}

func prepareScenarios() map[io.Reader][]Result {
	scenarios := make(map[io.Reader][]Result)
	scenarios[getFile("../testresources/ex1.html")] = []Result{{link: "/other-page", text: "A link to another page"}}
	return scenarios
}

func getFile(fileName string) io.Reader {
	file, error := os.Open(fileName)
	if error != nil {
		log.Fatalln("Cannot open file:", fileName)
	}
	return file
}
