package main

import "fmt"

type mechanic interface {
	execute(*car)
	setNext(mechanic)
}

type fuelbag struct {
	next mechanic
}

func (r *fuelbag) execute(p *car) {
	if p.Fuel {
		fmt.Println("fuel in fuelbag")
		r.next.execute(p)
		return
	}
	fmt.Println("fuel ready!")
	p.Fuel = true
	r.next.execute(p)
}

func (r *fuelbag) setNext(next mechanic) {
	r.next = next
}

type battery struct {
	next mechanic
}

func (r *battery) execute(p *car) {
	if p.Energy {
		fmt.Println("power alreay on")
		r.next.execute(p)
		return
	}
	fmt.Println("power on!")
	p.Energy = true
	r.next.execute(p)
}

func (r *battery) setNext(next mechanic) {
	r.next = next
}

type keycar struct {
	next mechanic
}

func (r *keycar) execute(p *car) {
	if p.key {
		fmt.Println("key already in car")
		r.next.execute(p)
		return
	}
	fmt.Println("key in car")
	p.key = true
	r.next.execute(p)
}

func (r *keycar) setNext(next mechanic) {
	r.next = next
}

type car struct {
	name   string
	Fuel   bool
	Energy bool
	key    bool
	build  bool
}

type Engine struct {
	next mechanic
}

func (c *Engine) execute(p *car) {
	if p.build {
		fmt.Println("already done")
	}
	fmt.Println("have make systeam to drive")
}

func (c *Engine) setNext(next mechanic) {
	c.next = next
}

func main() {
	Eng := &Engine{}
	work := &fuelbag{}
	work.setNext(Eng)
	//
	work2 := &battery{}
	work2.setNext(work)
	//
	work3 := &keycar{}
	work3.setNext(work2)
	//
	car2 := &car{name: "HONDA"}
	work3.execute(car2)
}
