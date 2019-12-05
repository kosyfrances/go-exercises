package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	err := tmpl.Execute(w, h.adv["intro"])
	if err != nil {
		fmt.Println(err)
	}
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
