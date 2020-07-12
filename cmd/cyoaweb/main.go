package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Loading %s\n", *filename)

	f,err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n",story)
}
