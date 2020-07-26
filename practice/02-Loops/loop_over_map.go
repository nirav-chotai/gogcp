// Activity 2.02

package main

import (
	"fmt"
)

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	highestCount := 0
	finalWord := ""

	for word, count := range words {
		if count > highestCount {
			highestCount = count
			finalWord = word
		}
	}

	fmt.Println("Most popular word: ", finalWord)
	fmt.Println("With a count of : ", highestCount)
}
