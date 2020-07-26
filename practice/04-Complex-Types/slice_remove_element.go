// Activity 4.05

package main

import "fmt"

func main() {
	outcome := []string{"Good", "Good", "Bad", "Good", "Good"}
	finalSlice := []string{}
	finalSlice = append(finalSlice, outcome[0:2]...)
	finalSlice = append(finalSlice, outcome[3:5]...)
	fmt.Println(finalSlice)
}
