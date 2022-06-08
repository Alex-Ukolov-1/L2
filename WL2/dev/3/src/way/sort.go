package ways

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TextFile struct {
	file      string
	sort      bool
	filend    string
	operation string
	dateStr   []string
}

func NeuTextFile() *TextFile {
	fmt.Print("Введите файл начальный: ")
	var input string
	fmt.Fscan(os.Stdin, &input)

	fmt.Print("Введите файл конечный: ")
	var input2 string
	fmt.Fscan(os.Stdin, &input2)

	fmt.Print("Введите операцию k/n/r/u: ")
	var operation string
	fmt.Fscan(os.Stdin, &operation)

	return &TextFile{
		file:      input,
		sort:      false,
		filend:    input2,
		operation: operation,
	}

}

//os.Args[0]

func (text *TextFile) Read() {
	fmt.Println(" ")
	file, err := os.Open(text.file)

	if err != nil {
		log.Fatal("open file error: ", err)
	}
	defer file.Close()

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("read file error: ", err)
	}

	text.dateStr = strings.Split(string(dataBytes), "\n")
}

func (text *TextFile) Write() {

	file, err := os.Create(text.filend)
	if err != nil {
		log.Fatal("create file error: ", err)
	}

	defer file.Close()

	v := selectoperation(text.dateStr, text.operation)

	for _, line := range v {
		if _, err = file.WriteString(line + "\n"); err != nil {
			log.Fatal("write file error: ", err)
		}
	}
	fmt.Println(v)
}

func selectoperation(file []string, operation string) []string {
	var v []string
	switch {
	case operation == "k":
		v = Columnndefinite(file, 0, true)
	case operation == "n":
		v = sortbynumber(file, true)
	case operation == "r":
		v = sortback(file)
	case operation == "u":
		v = donotshowrepeatstring(file)
	default:
		fmt.Println("Invalid")
	}
	return v
}

func Columnndefinite(file []string, k int, n bool) []string {
	s := make([][]string, 0)

	k = k - 1
	if k < 0 {
		k = 0
	}

	for _, line := range file {
		s = append(s, strings.Split(line, " "))
	}

	if n {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				x, err := strconv.Atoi(s[i][k])
				y, err := strconv.Atoi(s[j][k])
				if err != nil {
					fmt.Println(err)
					return false
				}

				return x < y
			}

			return false
		})
	} else {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				return strings.ToLower(s[i][k]) < strings.ToLower(s[j][k])
			}
			return false
		})
	}

	var str string
	sl := make([]string, 0)
	// строка файла которая была разделена пробелом, джониться обратно пробелом
	for _, line := range s {
		str = strings.Join(line, " ")
		sl = append(sl, str)
	}

	// возвращаем уже отсортированный слайс
	return sl
}

func sortbynumber(file []string, numeric bool) []string {
	if numeric == true {
		numbers := []float64{}
		strs := []string{}
		for _, s := range file {
			if isNumeric(s) == true {
				n, _ := strconv.ParseFloat(s, 64)
				numbers = append(numbers, n)
			} else {
				strs = append(strs, s)
			}
		}
		sort.Strings(strs)
		sort.Float64s(numbers)
		for _, f := range numbers {
			s := strconv.FormatFloat(f, 'f', -1, 64)
			strs = append(strs, s)
		}
		return strs
	} else {
		sort.Strings(file)
		return file
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func sortback(file []string) []string {
	for i, j := 0, len(file)-1; i < j; i, j = i+1, j-1 {
		file[i], file[j] = file[j], file[i]
	}

	// возвращаем уже отсортированный слайс
	return file
}

func donotshowrepeatstring(file []string) []string {
	data := make(map[string]struct{})
	for _, str := range file {
		if _, ok := data[str]; !ok {
			data[str] = struct{}{}
		}
	}
	resultData := make([]string, 0, len(data))
	for key := range data {
		resultData = append(resultData, key)
	}
	return resultData
}
