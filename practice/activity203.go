package main

import (
	"fmt"
)

func main() {

	unsortedNumber := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	sortedNumber := []int{8, 5, 2, 4, 0, 1, 3, 7, 9, 6}

	// number of items to be sorted
	n := len(unsortedNumber)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if sortedNumber[i] > sortedNumber[i+1] {
				sortedNumber[i], sortedNumber[i+1] = sortedNumber[i+1], sortedNumber[i]
				swapped = true
			}
		}
	}
	fmt.Println("Before: ", unsortedNumber)
	fmt.Println("After: ", sortedNumber)
}
