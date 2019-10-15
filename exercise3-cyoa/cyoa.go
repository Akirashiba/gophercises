package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"html/template"
)

type Option struct{
	Text string `json:"text"`
	Arc string `json:"arc"`
}

type StoryDetail struct{
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Option `json:"options"`
}

type Stories map[string]StoryDetail

var StorySets Stories

func (*Stories) handler(w http.ResponseWriter, r *http.Request){
	url := r.URL.String()
	if detail, ok := StorySets[url[1:]]; ok{
		tmpl := template.Must(template.ParseFiles("./layout.html"))
		tmpl.Execute(w, detail)
	}
}

func main(){
	err := getStoryMap()
	if err != nil{
		return
	}
	mux := defaultMux()
	for k, _ := range StorySets{
		mux.HandleFunc(fmt.Sprintf("/%s", k), StorySets.handler)
	}
	fmt.Println("Starting the server on :8081")
	http.ListenAndServe(":8081", mux)
}

func getStoryMap()(err error){
	data, err := ioutil.ReadFile("./stories.json")
    if err != nil {
        return
	}
	err = json.Unmarshal(data, &StorySets)
    if err != nil {
        return
	}
	fmt.Println(StorySets)
	return
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}