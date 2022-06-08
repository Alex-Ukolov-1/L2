package main

import "fmt"

type Command interface {
	execute()
}

type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

func NewResteraunt() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "pizza")
}

type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "salad")
}

type CleanDishesCommand struct {
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {
	c.restaurant.CleanedDishes = c.restaurant.TotalDishes
	fmt.Println("dishes cleaned")
}

func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}

type Cook struct {
	Commands []Command
}

func (c *Cook) executeCommands() {
	for _, r := range c.Commands {
		r.execute()
	}
}

func main() {
	dish := NewResteraunt()
	work := []Command{dish.MakePizza(1), dish.MakeSalad(2), dish.MakePizza(3)}
	cooks := []Cook{}
	cooks = append(cooks, Cook{work})
	for _, c := range cooks {
		c.executeCommands()
	}

}
