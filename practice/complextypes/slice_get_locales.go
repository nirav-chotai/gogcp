package complextypes

import (
	"fmt"
	"os"
)

func getPassedArg() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func getLocals(extraLocals []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocals...)
	return locales
}

func sliceget() {
	locales := getLocals(getPassedArg())
	fmt.Println("Locales to use: ", locales)
}
