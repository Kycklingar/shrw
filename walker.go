package shrw

import (
	"golang.org/x/net/html"
)

// Get the first node matching m
func Walk(node *html.Node, m Matcher) *html.Node {
	if n := m.Match(node); n != nil {
		return n
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if n := Walk(c, m); n != nil {
			return n
		}

	}

	return nil
}

// Get all nodes matching m
func WalkAll(ch chan *html.Node, node *html.Node, m Matcher) {
	var f func(node *html.Node, m Matcher)

	f = func(node *html.Node, m Matcher) {
		if n := m.Match(node); n != nil {
			ch <- n
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c, m)
		}
	}

	f(node, m)
	close(ch)
}

// Get first grandchild node matching the m pattern
func WalkPattern(node *html.Node, m ...Matcher) *html.Node {
	return walkPattern(node, 0, m...)
}

func walkPattern(node *html.Node, i int, m ...Matcher) *html.Node {
	if n := m[i].Match(node); n != nil {
		i++
		if i >= len(m) {
			return n
		}
	} else {
		i = 0
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if n := walkPattern(c, i, m...); n != nil {
			return n
		}
	}

	return nil

}

// Get all grandchild nodes matching the m pattern
func WalkPatternAll(ch chan *html.Node, node *html.Node, m ...Matcher) {
	var f func(*html.Node, int, ...Matcher)
	f = func(node *html.Node, i int, m ...Matcher) {
		if n := m[i].Match(node); n != nil {
			i++
			if i >= len(m) {
				ch <- n
				i = 0
			}
		} else {
			i = 0
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c, i, m...)
		}
	}

	f(node, 0, m...)
	close(ch)
}
