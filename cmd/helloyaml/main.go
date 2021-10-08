package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Response struct {
	Title    string   `yaml:"title"`
	Messages []string `yaml:"messages"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Title:    "Some Title",
		Messages: []string{"First message", "second message"},
	}
	respYAML, err := yaml.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(respYAML))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
