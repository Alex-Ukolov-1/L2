package main

import (
	"log"
)

func main() {
	product := "car"
	payWay := 3

	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("831", "456831904")
	case 2:
		payment = NewPayPalPayment("1443566")
	case 3:
		payment = NewQIWIPayment("1", "02.06.2022", "Moscow")
	}

	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	payment.Pay()
}

type Payment interface {
	Pay()
}

type cardPayment struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() {
	log.Println("Sucsesfull cardpayment!")
}

type payPalPayment struct {
	cardNumber string
}

func NewPayPalPayment(cardNum string) *payPalPayment {
	return &payPalPayment{cardNumber: cardNum}
}

func (p *payPalPayment) Pay() {
	log.Println("Sucsesfull payPalPayment!")
}

type qiwiPayment struct {
	cardNumber, cvv, city string
}

func NewQIWIPayment(cardNum, cvv, city string) Payment {
	return &qiwiPayment{
		cardNumber: cardNum,
		cvv:        cvv,
		city:       city,
	}
}

func (p *qiwiPayment) Pay() {
	log.Println("Sucsesfull qiwipayment!")
}
