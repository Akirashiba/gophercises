package main

import (
	"fmt"
	"net/http"
	"encoding/xml"
	"flag"
	link "gophercises/exercise4-linkparser"
)

type linkList []*link.LinkNode

var ReachedPage []string 

func GetPageLinkList(pageUrl string)(linklist linkList, err error){
	resp, err := http.Get(pageUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	return	
}

func Go

func main() {
	
}