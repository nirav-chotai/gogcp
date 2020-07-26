package main

import "fmt"

func main() {
	itemSold()
}

func itemSold() {
	items := make(map[string]int)
	items["Nirav"] = 41
	items["Gunja"] = 109
	items["Harsh"] = 24

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)
		if v < 40 {
			fmt.Println("is below expectations.")
		} else if v >= 40 && v <= 100 {
			fmt.Println("meet expectations.")
		} else if v > 100 {
			fmt.Println("above expecations.")
		}
	}
}
