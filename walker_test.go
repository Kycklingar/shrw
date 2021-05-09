package shrw

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const htmlString = `
<html>
<head>
</head>
<body>
	<div class="findme">
		<div class="find me too">
			<p>Hello <a id="thea">World</a></p>
		</div>
	</div>
	<div>
		<div>
			<a>Hello</a>
		</div>
	</div>
	<span>
		Text
		<div>hello</div>
	</span>
</body>
</html>
`

func TestWalk(t *testing.T) {
	r := strings.NewReader(htmlString)

	n, err := html.Parse(r)
	if err != nil {
		t.Fatal(err)
	}

	findme := Walk(n, ClassAll("findme"))
	if findme == nil {
		t.Fatal("findme is nil")
	}

	findto := Walk(n, Class("me"))
	if findto == nil {
		t.Fatal("findto is nil")
	}

	findid := Walk(n, Id("thea"))
	if findid == nil {
		t.Fatal("findid is nil")
	}

	findp := Walk(n, Tag("p"))
	if findp == nil {
		t.Fatal("findp is nil")
	}

	findText := Walk(n, Text("World"))
	if findText == nil {
		t.Fatal("findText is nil")
	}

	findpattern := WalkPattern(n, 0, Tag("div"), Tag("div"), Tag("a"))
	if findpattern == nil {
		t.Fatal("pattern is nil")
	}

	if findpattern.FirstChild.Data != "Hello" {
		t.Fatal("findpattern data incorrect. Expected: Hello, Got: ", findpattern.FirstChild.Data)
	}

	findtextblock := WalkPattern(n, 0, Tag("span"), Tag("div"))
	if findtextblock == nil {
		t.Fatal("findtextblock is nil")
	}

	if findtextblock.FirstChild.Data != "hello" {
		t.Fatal("expect hello got ", findtextblock.FirstChild.Data)
	}

}

func TestWalkAll(t *testing.T) {
	r := strings.NewReader(htmlString)

	n, err := html.Parse(r)
	if err != nil {
		t.Fatal(err)
	}

	ch := make(chan *html.Node, 2)
	go WalkAll(ch, n, Tag("a"))
	world, hello := <-ch, <-ch
	if world.FirstChild.Data != "World" || hello.FirstChild.Data != "Hello" {
		t.Fatal("Did not get hello world")
	}

	if _, ok := <-ch; ok {
		t.Fatal("Got to many nodes")
	}

	ch = make(chan *html.Node, 2)
	go WalkPatternAll(ch, n, Tag("div"), Tag("div"), Tag("a"))
	var i int
	for n := range ch {
		i++
		if n.FirstChild.Data != "Hello" {
			t.Fatal("Expected Hello got ", n.FirstChild.Data)
		}
	}

	if i > 1 {
		t.Fatal("To many nodes")
	}
}
