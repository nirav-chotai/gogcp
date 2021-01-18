// Activity 4.01

package complextypes

import (
	"fmt"
)

func myarray() {
	var name [10]int
	for i := 0; i < len(name); i++ {
		name[i] = i + 1
	}
	fmt.Println(name)
}
