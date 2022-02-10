package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	cyoa "github.com/hazunanafaru/gophercises-cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON story file you will use")
	flag.Parse()
	fmt.Printf("Using the story in %s.", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
