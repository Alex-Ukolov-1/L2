package main

import (
	"fmt"
)

type Housepart interface {
	Accept(HousePartVisitor)
}

type Room struct {
	Name string
}

func (this *Room) Accept(visitor HousePartVisitor) {
	visitor.visitRoom(this)
}

type Part struct{}

func (this *Part) Accept(visitor HousePartVisitor) {
	visitor.visitHouse(this)
}

type House struct {
	parts []Housepart
}

func NewHouse() *House {
	this := new(House)
	this.parts = []Housepart{
		&Room{"first room"},
		&Room{"second room"},
		&Room{"third room"},
		&Room{"fourth room"},
		&Part{}}
	return this
}

func (this *House) Accept(visitor HousePartVisitor) {
	for _, part := range this.parts {
		part.Accept(visitor)
	}
}

type HousePartVisitor interface {
	visitRoom(room *Room)
	visitHouse(house *Part)
}

type GetMessageVisitor struct {
	Messages []string
}

func (this *GetMessageVisitor) visitRoom(Room *Room) {
	this.Messages = append(this.Messages, fmt.Sprintf("Visiting the %v\n", Room.Name))
}

func (this *GetMessageVisitor) visitHouse(House *Part) {
	this.Messages = append(this.Messages, fmt.Sprintf("Visiting HOUSE\n"))
}

func main() {
	one := NewHouse()
	visitor := new(GetMessageVisitor)
	one.Accept(visitor)
	fmt.Println(visitor.Messages)
}
