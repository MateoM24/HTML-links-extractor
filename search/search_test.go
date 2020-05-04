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
				if expected.Url == got.Url {
					foundLink = true
					if expected.Text == got.Text {
						break
					} else {
						t.Errorf("Didn't find record having text: [%s]. Got instead: [%s]", expected.Text, got.Text)
					}
				}
			}
			if !foundLink {
				t.Errorf("Didn't find record having link: [%s].", expected.Url)
			}
		}
	}
}

func prepareScenarios() map[io.Reader][]Link {
	scenarios := make(map[io.Reader][]Link)
	scenarios[getFile("../testresources/ex1.html")] = []Link{{Url: "/other-page", Text: "A link to another page"}}
	scenarios[getFile("../testresources/ex2.html")] = []Link{
		{Url: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
		{Url: "https://github.com/gophercises", Text: "Gophercises is on Github"},
	}
	scenarios[getFile("../testresources/ex3.html")] = []Link{
		{Url: "#", Text: "Login"},
		{Url: "/lost", Text: "Lost? Need help?"},
		{Url: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
	}
	scenarios[getFile("../testresources/ex4.html")] = []Link{{Url: "/dog-cat", Text: "dog cat"}}
	return scenarios
}

func getFile(fileName string) io.Reader {
	file, error := os.Open(fileName)
	if error != nil {
		log.Fatalln("Cannot open file:", fileName)
	}
	return file
}
