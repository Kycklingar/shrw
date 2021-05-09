package shrw

import (
	"strings"

	"golang.org/x/net/html"
)

type Matcher interface {
	// Match returns the node if it matches
	// it will otherwise return nil
	Match(*html.Node) *html.Node
}

// Implements the matcher interface
type (
	Id         string
	Class      string // "one" or "two" will match class="one two"
	ClassFull  string // only "one two" will match class="one two"
	Tag        string
	Text       string
	TextNoTrim string
)

func (m Id) Match(node *html.Node) *html.Node {
	return checkAttribute(node, "id", string(m))
}

func (m Class) Match(node *html.Node) *html.Node {
	for _, attr := range node.Attr {
		if attr.Key == "class" {
			for _, v := range strings.Split(attr.Val, " ") {
				if v == string(m) {
					return node
				}
			}
		}
	}

	return nil
}

func (m ClassFull) Match(node *html.Node) *html.Node {
	return checkAttribute(node, "class", string(m))
}

func (m Tag) Match(node *html.Node) *html.Node {
	if node.Type == html.ElementNode && node.Data == string(m) {
		return node
	}

	return nil
}

func (m Text) Match(node *html.Node) *html.Node {
	if node.Type == html.TextNode && strings.TrimSpace(node.Data) == string(m) {
		return node
	}

	return nil
}

func (m TextNoTrim) Match(node *html.Node) *html.Node {
	if node.Type == html.TextNode && node.Data == string(m) {
		return node
	}

	return nil
}

func checkAttribute(node *html.Node, k, v string) *html.Node {
	for _, attr := range node.Attr {
		if attr.Key == k && attr.Val == v {
			return node
		}
	}

	return nil
}
