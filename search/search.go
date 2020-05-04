package search

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func RetrieveLinks(file io.Reader) *[]Link {
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalln("Cannot parse HTML file", err, file)
	}
	links := make([]Link, 0, 10)
	findLinkInNodes(doc, &links)
	return &links
}

func findLinkInNodes(n *html.Node, results *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, att := range n.Attr {
			if att.Key == "href" {
				finding := Link{}
				finding.Url = att.Val
				finding.Text = findText(n)
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
		sibling := n.NextSibling
		if sibling != nil {
			return strings.TrimSpace(strings.Join([]string{strings.TrimSpace(n.Data), findText(sibling)}, " "))
		}
		return strings.TrimSpace(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return findText(c)
	}
	return ""
}

type Link struct {
	Url, Text string
}
