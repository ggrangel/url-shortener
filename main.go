package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ggrangel/url-shortener/urlshort"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v3",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on :8080")
	if err := http.ListenAndServe(":8080", yamlHandler); err != nil {
		log.Fatal(err)
	}
}
