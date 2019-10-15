package link

import (
	"fmt"
	"golang.org/x/net/html"
)

type LinkNode struct {
	Href string
	Text string
}


func ParseLink(linkList *[]*LinkNode, doc *html.Node)(err error){
	var f func(n *html.Node)
	fmt.Println("linkList", linkList)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			href := n.Attr[0].Val 
			text := n.FirstChild.Data
			newNode := &LinkNode{href, text}
			*linkList = append(*linkList, newNode)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil
}



