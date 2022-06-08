package main

import (
	"log"
)

type Bank struct {
	id        int
	name      string
	operation string
}

type Connect interface {
	setid(id int) Connect
	setname(name string) Connect
	setoperation(operation string) Connect
}

func newConnection() Connect {
	return &Bank{
		id:        0,
		name:      "",
		operation: "",
	}
}

func (h Bank) setid(id int) Connect {
	h.id = id
	return h
}

func (h Bank) setname(name string) Connect {
	h.name = name
	return h
}

func (h Bank) setoperation(operation string) Connect {
	h.operation = operation
	return h
}

func main() {
	operation := newConnection().setid(1).setname("Sberbank").setoperation("Spisaniye")
	log.Println(operation)
	operation1 := newConnection().setid(2).setname("Tinkoff").setoperation("Zachisleniye")
	log.Println(operation1)

}
