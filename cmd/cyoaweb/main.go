package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "port to start the web application")
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Loading %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		log.Fatal(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("Staring server on at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
