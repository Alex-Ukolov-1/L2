package main

import (
	"anagram"
	"fmt"
	//"sort"
	//"os"
	//"strconv"
)

func main() {
	var anagrams = []string{"пятак", "пятка", "пятка", "тяпка", "листок", "слиток", "столик", "слиток"}
	an := anagram.Findanagrams(&anagrams)
	for key, val := range an {
		fmt.Println(key, val)
	}
}
