package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	const fileName = "ex1.html"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Cannot open file:", fileName)
	}
	// bytes, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatalln("Cannot read file:", file)
	// }
	// fmt.Println(string(bytes))

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalln("Cannot parse HTML file:", fileName)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, att := range n.Attr {
				fmt.Println(att.Key, att.Val)
			}
			fmt.Println()
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}
