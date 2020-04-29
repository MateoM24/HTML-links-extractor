package search

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

func RetrieveLinks(file io.Reader) *[]Result {
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalln("Cannot parse HTML file", err, file)
	}
	results := make([]Result, 0, 10)

	findLinkInNodes(doc, &results)
	return &results

}

func findLinkInNodes(n *html.Node, results *[]Result) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, att := range n.Attr {
			if att.Key == "href" {
				finding := Result{}
				finding.link = att.Val
				finding.text = findText(n)
				*results = append(*results, finding)
			}

		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLinkInNodes(c, results)
	}
}

func findText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return findText(c)
	}
	return ""
}

type Result struct {
	link, text string
}