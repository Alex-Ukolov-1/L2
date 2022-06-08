package main

import (
	"fmt"
	//"os"
	"strconv"
	"strings"
	"unicode"
)

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func check(name string) {

	if len(name) <= 0 {
		fmt.Println(" ")
	} else if IsLetter(name) == true {
		fmt.Println(name)
	} else if _, err := strconv.Atoi(name); err == nil {
		fmt.Printf("%q looks like a number.\n", name)
	} else {
		doit(name)
	}
}

func doit(name string) {
	var chislo []string
	var position []int
	var position2 []int

	ss := strings.Split(name, "")

	//преобразование символов для поиска чисел после букв
	//ковертация
	for i := 0; i < len(ss); i++ {
		aa, _ := strconv.Atoi(ss[i])

		for j := 0; j < len(ss); j++ {

			if ss[j] == strconv.Itoa(aa) {
				chislo = append(chislo, ss[j])
			}

		}
		if (ss[i]) != strconv.Itoa(aa) {
			position = append(position, i)
		}
	}

	//fmt.Println(chislo)

	//fmt.Println(position)
	//алгоритм для нахождения позиций одиночных символов
	//
	for i := 0; i < len(position)-1; i++ {

		if position[i+1]-position[i] == 1 {
			position2 = append(position2, position[i])
		} else if position[i] == len(position) {
			position2 = append(position2, position[i+1])
		}

	}

	//fmt.Println(position2)
	//алгоритм для повторения букв
	//в том числе и вывод одиночных символов
	var answer []string

	for i := 0; i < len(chislo); i++ {
		for j := 0; j < len(ss); j++ {
			aa, _ := strconv.Atoi(chislo[i])

			for k := 0; k < aa; k++ {

				if chislo[i] == ss[j] {
					answer = append(answer, ss[j-1])
				}

				for l := 0; l < len(position2)-1; l++ {

					if j == position2[l] {

						answer = append(answer, ss[j])

						position2 = append(position2[:l], position2[l+1:]...)
					}

				}
			}
		}
	}

	fmt.Println(answer)
}

func main() {

	//var name string
	//fmt.Println("Введите строку: ")
	//fmt.Fscan(os.Stdin, &name)
	check("qwe\\5")
	check("a4bc2d5e3")
}
