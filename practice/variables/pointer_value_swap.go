// Activity 1.02

package variables

import "fmt"

func myswap() {
	a, b := 5, 10
	// call swap here
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}
func swap(a *int, b *int) {
	// swap the values here
	*a, *b = *b, *a
}
