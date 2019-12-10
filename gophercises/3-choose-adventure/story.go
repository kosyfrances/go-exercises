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

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

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

type handlerOption func(h *handler)

type handler struct {
	adv    adventure
	t      *template.Template
	pathFn func(r *http.Request) string
}

func withTemplate(t *template.Template) handlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

func withPathFunc(fn func(r *http.Request) string) handlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func newHandler(adv adventure, opts ...handlerOption) http.Handler {
	h := handler{adv, tmpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

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

func mapAdventure(content []byte) (adventure, error) {
	var adv adventure
	err := json.Unmarshal(content, &adv)
	if err != nil {
		return nil, err
	}

	return adv, nil
}
