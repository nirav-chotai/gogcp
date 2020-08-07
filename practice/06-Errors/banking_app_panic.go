package main

import (
	"errors"
	"fmt"
	"strings"
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func (dd *directDeposit) validateRoutingNumber() {
	if dd.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
}

func (dd *directDeposit) validateLastName() error {
	dd.lastName = strings.TrimSpace(dd.lastName)
	if len(dd.lastName) == 0 {
		return ErrInvalidLastName
	}
	return nil
}

func (dd *directDeposit) report() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Last name: ", dd.lastName)
	fmt.Println("First name: ", dd.firstName)
	fmt.Println("Bank name: ", dd.bankName)
	fmt.Println("Routing number: ", dd.routingNumber)
	fmt.Println("Account number: ", dd.accountNumber)
}

func main() {
	dd := directDeposit{
		lastName:      "",
		firstName:     "Nirav",
		bankName:      "ABC",
		routingNumber: 17,
		accountNumber: 1809,
	}

	dd.validateRoutingNumber()
	err := dd.validateLastName()
	if err != nil {
		fmt.Println(err)
	}
	dd.report()
}
