package main

import (
	"fmt"
	"strings"
)

func greeting() {
	fmt.Println("Please enter your name")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi, %s! I'm Go!", name)
	name = strings.TrimSpace(name)
}
