package main

import (
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

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
