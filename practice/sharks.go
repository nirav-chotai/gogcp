package practice

import (
	"fmt"
)

func sharks() {
	// Define sharks variable as a slice of strings
	sharks := []string{"hammerhead", "great white", "dogfish"}

	// For loop that iterates over sharks list and print each string
	for _, shark := range sharks {
		fmt.Println(shark)
	}
}
