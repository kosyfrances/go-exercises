package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := flag.String("file", "gopher.json", "the json adventure story file")
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

	fmt.Printf("%+v\n", adv)
}
