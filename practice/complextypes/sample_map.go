// Activity 4.02

package complextypes

import (
	"fmt"
	"os"
)

func samplemap() {
	names := make(map[string]string)
	names["305"] = "Sue"
	names["204"] = "Bob"
	names["631"] = "Jake"
	names["073"] = "Tracy"

	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("There was an error")
		os.Exit(1)
	}

	found := false
	for k := range names {
		if os.Args[1] == names[k] {
			fmt.Println("Hi, ", names[k])
			found = true
		}
	}

	if !found {
		fmt.Println("Error not found")
	}
}
