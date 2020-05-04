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
		for _, expected := range e {
			var foundLink bool
			for _, got := range *results {
				if expected.Link == got.Link {
					foundLink = true
					if expected.Text == got.Text {
						break
					} else {
						t.Errorf("Didn't find record having text: [%s]. Got instead: [%s]", expected.Text, got.Text)
					}
				}
			}
			if !foundLink {
				t.Errorf("Didn't find record having link: [%s].", expected.Link)
			}
		}
	}
}

func prepareScenarios() map[io.Reader][]Result {
	scenarios := make(map[io.Reader][]Result)
	scenarios[getFile("../testresources/ex1.html")] = []Result{{Link: "/other-page", Text: "A link to another page"}}
	scenarios[getFile("../testresources/ex2.html")] = []Result{
		{Link: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
		{Link: "https://github.com/gophercises", Text: "Gophercises is on Github"},
	}
	scenarios[getFile("../testresources/ex3.html")] = []Result{
		{Link: "#", Text: "Login"},
		{Link: "/lost", Text: "Lost? Need help?"},
		{Link: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
	}
	scenarios[getFile("../testresources/ex4.html")] = []Result{{Link: "/dog-cat", Text: "dog cat"}}
	return scenarios
}

func getFile(fileName string) io.Reader {
	file, error := os.Open(fileName)
	if error != nil {
		log.Fatalln("Cannot open file:", fileName)
	}
	return file
}
