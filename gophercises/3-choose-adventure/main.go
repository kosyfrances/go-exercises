package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the json adventure story file")
	port := flag.Int("port", 3000, "the port to start the server for the adventure story")
	flag.Parse()

	fmt.Printf("Using the adventure story file %v", *fileName)

	content, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	adv, err := mapAdventure(content)
	if err != nil {
		panic(err)
	}

	h := newHandler(adv)

	fmt.Printf("Starting server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
