package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the json adventure story file")
	port := flag.Int("port", 3000, "the port to start the server for the adventure story")
	flag.Parse()

	log.Printf("Using the adventure story file %v", *fileName)

	content, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	adv, err := mapAdventure(content)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.New("").Parse(storyTmpl))
	h := newHandler(adv, withTemplate(tmpl), withPathFunc(pathFn))

	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	log.Printf("Starting server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

var storyTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}
    <ul>
    {{range .Options}}
        <li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
    {{end}}
    </ul>
</body>
</html>
`
