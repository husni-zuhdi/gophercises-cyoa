package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON story file you will use")
	flag.Parse()
	fmt.Printf("Using the story in %s.", *filename)
}
