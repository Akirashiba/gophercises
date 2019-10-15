package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
	link "gophercises/exercise4-linkparser"
)

var fileLinkMap map[string][]*link.LinkNode

func main() {
	fileLinkMap = make(map[string][]*link.LinkNode)
	for i := 1; i <= 4; i++ {
		fileName := fmt.Sprintf("./ex%d.html", i)
		data, err := os.Open(fileName)
    	if err != nil {
			fmt.Println(err)
    	    return
		}
		doc, err := html.Parse(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		var linkList []*link.LinkNode
		link.ParseLink(&linkList, doc)
		fileLinkMap[fileName] = linkList
		fmt.Println(linkList)	
	}
	for file, linkMap := range fileLinkMap{
		fmt.Println("file:", file)
		for _, node := range linkMap{
			fmt.Println("a Node:")
			fmt.Println("    Text:", node.Text)
			fmt.Println("    Href:", node.Href)
		}
		
	}
}