package main

import "log"

type Animal interface {
	Sound() string
}

type Cat struct {
}

func NewCat() Cat {
	return Cat{}
}

func (c Cat) Sound() string {
	return "meow"
}

type Dog struct {
	name string
}

func NewDog() *Dog {
	return &Dog{name: "Doug"}
}

func (d Dog) Sound() string {
	return "woof"
}

func farm(x string) {
	var a Animal
	switch {
	case x == "cat":
		a = NewCat()
		log.Println(a.Sound())
	case x == "dog":
		a = NewDog()
		log.Println(a.Sound())
	default:
		log.Println("error")
	}
}

func main() {
	farm("cat")
	farm("dog")
}
