package practice

import (
	"fmt"
	"math/rand"
)

func random() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d) %d\n", i, rand.Intn(25))
	}
}
