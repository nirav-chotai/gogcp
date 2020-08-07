package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func main() {
	lastName := "chotai"
	routeNumber := 10

	if lastName == "chotai" {
		fmt.Println(ErrInvalidLastName)
	}
	if routeNumber == 10 {
		fmt.Println(ErrInvalidRoutingNumber)
	}
}
