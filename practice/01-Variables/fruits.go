package main

import (
	"fmt"
)

func fruits() {

	// Define fruits variable as a slice of strings
	fruits := []string{"Apple", "Orange", "Banana"}

	// For loop that iterate over fruits list and prints each string
	for _, fruits := range fruits {
		fmt.Println(fruits)
	}
}
