// Activity 4.04

package main

import "fmt"

func main() {
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	var newdays []string
	otherdays := weekdays[:6]
	newdays = append(newdays, weekdays[len(weekdays)-1])
	newdays = append(newdays, otherdays...)
	fmt.Println(newdays)
}
