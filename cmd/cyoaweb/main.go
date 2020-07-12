package main

import (
	"cyoa"
	"encoding/json"
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

	d := json.NewDecoder(f)

	var story cyoa.Story

	if err := d.Decode(&story); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n",story)
}
