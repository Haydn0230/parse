package Link

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// link represents the structure of the data returned
type link struct {
	Href string
	Text string
}
const (
	anchorTag = "a"
	strongTag = "strong"
)

// errCheck way to check errs
func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Parse reads a htmldoc and recursively goes through each node looking for a tags
func Parse(r io.Reader) ([]link, error){
	doc, err := html.Parse(r)
	errCheck(err)

	links := make([]link, 0)

	var parseTree func(*html.Node)

	// look for an 'a' tag
	parseTree = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == anchorTag {
			links = append(links, createLink(node))
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			parseTree(c)
		}
	}

	parseTree(doc)

	return links, nil
}

// createLink returns a parsed node in a link structure
func createLink(nodes *html.Node) link {
	for _, attribute := range nodes.Attr {
		return link{
			Href: attribute.Val,
			Text: strings.TrimSpace(parseText(nodes)),
		}
	}
	return link{}
}

// grabText traverses horizontally until next sibling is nil and down the tree when we hit a branch
func parseText(node *html.Node) string{
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}


	var parsedText string

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		parsedText = fmt.Sprintf("%s%s", parsedText, parseText(c))
	}

	return parsedText
}

