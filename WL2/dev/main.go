package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagvar string
	flag.StringVar(&flagvar, "name", "help message for flagname", "10")
	flag.Parse()

	fmt.Println("flagvar has value ", flagvar)
}
