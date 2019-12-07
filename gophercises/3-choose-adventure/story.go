package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var defaultHandlerTmpl = `
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
        <li><a href="/{{.Arc}}">{{.Text}}</a></li>
    {{end}}
    </ul>
</body>
</html>
`

type chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type adventure map[string]chapter

type handler struct {
	adv adventure
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Parse(defaultHandlerTmpl))

	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.adv[path]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Ooops! That chapter was not found.", http.StatusNotFound)
}

func newHandler(adv adventure) http.Handler {
	return handler{adv}
}

func mapAdventure(content []byte) (adventure, error) {
	var adv adventure
	err := json.Unmarshal(content, &adv)
	if err != nil {
		return nil, err
	}

	return adv, nil
}
