package main

import (
	"fmt"
	"time"
)

type Bank struct {
	amount int
}

type Client struct {
	amount int
}

type Facade struct {
	id int
	*Client
	*Bank
}

type CallCount interface {
	fixtime()
}

func (s Bank) fixtime() {
	fmt.Println(time.Now())
}

func (s Client) fixtime() {
	fmt.Println(time.Now())
}

func show(a int, b int, try bool) {
	if try == true {
		v := set(a, b)
		if v != nil {
			fmt.Println("Transaction open")
			v.Client.OpenTransaction()
			v.Bank.OpenTransaction()

			v.transferMoney(Client{}, Bank{})

			v.Client.CloseTransaction()
			v.Bank.CloseTransaction()
			fmt.Println("Transaction closed")
			fmt.Println(" ")
		} else {
			fmt.Println("error of server")
		}
	} else {
		fmt.Println("error of server")
	}
}

func (s *Bank) OpenTransaction() {
	fmt.Println("connection with bank")
}

func (s *Bank) CloseTransaction() {
	fmt.Println("closed with bank")
}

func (s *Facade) transferMoney(AA Client, BB Bank) {
	if s.Client.amount > s.Bank.amount {
		var online CallCount
		online = Client{}
		online.fixtime()
		abc := s.Client.amount - s.Bank.amount
		fmt.Println("balance", abc)
		online = Bank{}
		online.fixtime()
	} else {
		fmt.Println("not enought money")
	}
}

func (s *Client) OpenTransaction() {
	fmt.Println("connection with client")
}

func (s *Client) CloseTransaction() {
	fmt.Println("closed with client")
}

func set(a, b int) *Facade {
	return &Facade{1, &Client{amount: a}, &Bank{amount: b}}
}

func main() {
	arr1 := []int{1000, 2000, 3000, 4000, 5000}
	arr2 := []int{100, 200, 300, 400, 500}
	for i := 0; i < len(arr1); i++ {
		show(arr1[i], arr2[i], true)
	}
}
