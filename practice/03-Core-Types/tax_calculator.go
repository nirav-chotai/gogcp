// Activity 3.01

package main

import "fmt"

func salesTax(cost float64, rate float64) float64 {
	return (rate * cost) / 100
}

func main() {
	total := 0.00
	cakeCost := 0.99
	cakeTax := 7.5
	milkCost := 2.75
	milkTax := 1.5
	butterCost := 0.87
	butterTax := 2.0

	total = salesTax(cakeCost, cakeTax)
	total = total + salesTax(milkCost, milkTax)
	total = total + salesTax(butterCost, butterTax)

	fmt.Println("Sales Tax Total: ", total)

}
