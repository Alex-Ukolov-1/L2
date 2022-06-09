package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	glow, err := http.Get("https://yandex.ru/")
	if err != nil {
		log.Fatalf("error to get URL")
	}
	defer glow.Body.Close()

	output, err := os.Create("index.html")
	defer output.Close()

	io.Copy(output, response.Body)
}
