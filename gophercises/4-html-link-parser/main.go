package main

import (
	"fmt"
	"os"

	"github.com/kosyfrances/go-exercises/gophercises/4-html-link-parser/link"
)

func main() {
	file := "example-htmls/ex1.html"
	r, err := os.Open(file)

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer r.Close()

	links, err := link.Parse(r)
	if err != nil {
		fmt.Errorf("cannot parse links on %s", file)
	}
	fmt.Println(links)
}
