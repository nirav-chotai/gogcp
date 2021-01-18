// Activity 1.03

package variables

import "fmt"

func bug() {
	message := "anything"
	count := 5
	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not greater than 5"
	}
	fmt.Println(message)
}
