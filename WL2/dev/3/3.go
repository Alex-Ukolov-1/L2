package main

import (
	//"fmt"
	"way"
)

//os.Args[0]

func main() {
	data := ways.NeuTextFile()
	data.Read()
	data.Write()
}
