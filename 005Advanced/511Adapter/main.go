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

//Creating adapter
type BankPaymetAdater struct {
	bankAccount int
	BankPayment *BankPayment
}

//Using the same interface
func (bpa BankPaymetAdater) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

//Main
func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	// cash.Pay()
	bpa := &BankPaymetAdater{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
