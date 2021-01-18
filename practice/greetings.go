package practice

import (
	"fmt"
	"strings"
)

func greetings() {
	fmt.Println("Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi, %s! I'm Go!", name)
	name = strings.TrimSpace(name)
}
