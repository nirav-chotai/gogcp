// Activity 1.04

package variables

import "fmt"

func count() {
	count := 0
	if count < 5 {
		count = 10
		count++
	}
	fmt.Println(count == 11)
}
