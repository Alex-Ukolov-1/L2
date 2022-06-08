package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://vk.com/id228187696")
	if err != nil {
		log.Fatalf("error to get URL")
	}
	defer response.Body.Close()

	output, err := os.Create("index.html")
	defer output.Close()

	io.Copy(output, response.Body)
}
