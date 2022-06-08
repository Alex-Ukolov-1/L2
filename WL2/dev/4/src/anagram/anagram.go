package anagram

import (
	//"fmt"
	//"os"
	"sort"
	//"strconv"
	"strings"
)

type AnagramStruct struct {
}

func newAnagramStruct() *AnagramStruct {
	return &AnagramStruct{}
}

func (v *AnagramStruct) toLowerCase(a string) string {
	s := strings.ToLower(a)
	return s
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func AnagramDict(in []string) map[string][]string {
	tempM := make(map[string][]string, 0) //промежуточная мапа, ключ - отсортированное слов

	for _, v := range in {
		sorted := []rune(v)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})
		sortedS := string(sorted) //отсортированное слово
		tempM[sortedS] = append(tempM[sortedS], v)
	}
	//результирующая мапа
	resultM := make(map[string][]string, 0)
	for _, v := range tempM {
		if len(v) > 1 { //если всего один элемент - в словрь не попадает
			resultM[v[0]] = v //нулевой элемент, это первый добавленный (первый просмотренный)
			sort.Strings(v)
		}
	}
	return resultM
}

func Uniq(set map[int][]string) map[string][]string {
	v := set[1]
	//fmt.Println(v)
	cc := removeDuplicateStr(v)
	vv := AnagramDict(cc)
	//fmt.Println(vv)
	//сюда!
	return vv
}

func AddElement(set map[int][]string, element string) {
	var key int
	key = key + 1
	value := set[key]
	set[key] = append(value, element)
}

func Findanagrams(array *[]string) map[string][]string {
	v := newAnagramStruct()
	for _, word := range *array {
		v.toLowerCase(word)
	}

	set := make(map[int][]string)
	for _, word := range *array {
		AddElement(set, word)
	}
	volga := Uniq(set)

	return volga
}
