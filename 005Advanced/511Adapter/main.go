package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

//Cashpayment implementa la interface

func (CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct {
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using Bankaccount %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	BankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.BankAccount)
}

//Main
func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	cash.Pay()
	bpa := &BankPaymentAdapter{
		BankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
