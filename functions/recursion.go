package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var webSite = `
<html>
  <head>
    <title>your title here</title>
  </head>
  <body bgcolor="ffffff">
    <center>
      <img src="clouds.jpg" align="bottom">
    </center>
    <hr>
    <a href="http://somegreatsite.com">link name</a>
    is a link to another nifty site
    <h1>this is a header</h1>
    <h2>this is a medium header</h2>
    <a href="http://text2.com">link name</a>
    send me mail at <a href="mailto:support@yourcompany.com">support@yourcompany.com</a>
    <p> this is a new paragraph! </p>
    <p>
      <b>this is a new paragraph!</b>
      <br> <b><i>this is a new sentence without a paragraph break, in bold italics.</i></b>
      <hr>
    </p>
  </body>
</html>`

func searchHtml() {
	doc, err := html.Parse(strings.NewReader(webSite))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	//outline(nil, doc)
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && len(n.Parent.Data) > 0{
		fmt.Printf("%s.%s -> ",n.Parent.Data, n.Data)
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
