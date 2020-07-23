package main

import (
	"fmt"
)

func main() {
	n := 1
	r := 100
	for n < r {
		if n%3 == 0 && n%5 == 0 {
			fmt.Println("FizzBuzz.")
		} else if n%3 == 0 {
			fmt.Println("Fizz.")
		} else if n%5 == 0 {
			fmt.Println("Buzz.")
		} else {
			fmt.Println(n)
		}
		n++
	}
}
