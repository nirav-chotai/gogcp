package main

import (
    "fmt"
    "github.com/nirav-chotai/gogcp/greet"
)

func main() {
    greet.Hello()

    fmt.Println(greet.Shark)

    oct := greet.Octopus{
        Name: "Nirav",
        Color: "Green",
    }

    fmt.Println(oct.String())
}