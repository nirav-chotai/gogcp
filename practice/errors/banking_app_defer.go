package errors

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
	errInvalidLastName      = errors.New("invalid last name")
	errInvalidRoutingNumber = errors.New("invalid routing number")
)

func (dd *directDeposit) validateRoutingNumber() {
	defer func() {
		r := recover()
		if r == ErrInvalidRoutingNumber {
			fmt.Println("invalid routing number")
		}
	}()
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

func bankapp() {
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
