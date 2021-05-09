# Simple Recursive Html Walker

Using a simple matching interface to find nodes in a *html.Node tree

```golang
interface Matcher {
	// Match returns the node if it matches
	// it will otherwise return nil
	Match(*html.Node) *html.Node
}

// Implements the Matcher interface
type (
	Id		string
	Class		string	// "one" or "two" will match class="one two"
	ClassFull	string	// only "one two" will match class="one two"
	Tag		string
	Text		string
	TextNoTrim	string
)
```
There are 4 flavors of walkers
```golang
// Get the first node matching
func Walk(*html.Node, Matcher)

// Get all nodes matching
func WalkAll(chan *html.Node, *html.Node, Matcher)

// Get first grandchild node matching the pattern
func WalkPattern(*html.Node, ...Matcher)

// Get all grandchild nodes matching the pattern
func WalkPatternAll(chan *html.Node, *html.Node, ...Matcher)
```
## Examples
Get node with id "horse"
```golang
node = shrw.Walk(node, shrw.Id("horse"))
```

Get all nodes with the "thumbnail" class
```golang
ch := make(chan *html.Node)

go shrw.WalkAll(ch, node, shrw.Class("thumbnail"))

for node := range ch {
	// Use node
}
```

Get node following a pattern
```golang
// This will get you the text node 'This Title' of
//	<div id="section">
//		<div class="message>
//			<h1>
//				This Title
//			</h1>
//		</div>
//	</div>
node = shrw.WalkPattern(
		node,
		shrw.Id("section"),
		shrw.Class("message"),
		shrw.Tag("h1"),
		shrw.Text("This Title"),
)
```
