package main

import (
	"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
	"bytes"
	"os"
)

type urlMap struct {
	paths map[string]string
}

var globalMap *urlMap = &urlMap{make(map[string]string)}

func (*urlMap) redirect(w http.ResponseWriter, r *http.Request) {
	if url, ok := globalMap.paths[r.URL.String()]; ok {
		http.Redirect(w, r, url, 307)
		return
	}
}

func MapHandler(pathsToUrls map[string]string, fallback *http.ServeMux) http.HandlerFunc {
	for k, v := range pathsToUrls{
		globalMap.paths[k] = v
		fallback.HandleFunc(k, globalMap.redirect)
	}
	return nil
}

func YAMLHandler(yaml []byte, fallback *http.ServeMux) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
  	if err != nil {
  	  	return nil, err
  	}
  	pathMap := buildMap(parsedYaml)
  	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yamlbytes []byte) (parsedYaml []map[string]string, err error){
	decoder := yaml.NewDecoder(bytes.NewBuffer(yamlbytes))
	err = decoder.Decode(&parsedYaml)
	if err != nil{
		fmt.Println(err)
		os.Exit(-2)
	}
	fmt.Println(parsedYaml)
	return
}

func buildMap(parsedYaml []map[string]string)(pathMap map[string]string){
	pathMap = make(map[string]string)
	for _, ref := range parsedYaml{
		fmt.Println(ref["url"])
		path, ok := ref["path"]
		if !ok {
			continue
		}
		url, ok := ref["url"]
		if !ok {
			continue
		}
		pathMap[path] = url
	}
	return
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	_, err := YAMLHandler([]byte(yaml), mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8081", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}