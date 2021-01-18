package functions

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var helloList = []string{
	"Nirav Chotai",
	"Gunja Chotai",
	"Nikunj Chotai",
	"Mayuri Chotai",
	"Henil Chotai",
}

func random() {
	// Seed random number generator using the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number in the range of our list
	index := rand.Intn(len(helloList))

	// Call a function and receive multiple return values
	msg, err := hello(index)

	// Handle any errors
	if err != nil {
		log.Fatal(err)
	}

	// Print our message to the console
	fmt.Println(msg)

	mymsg, myerr := star(index)
	if myerr != nil {
		log.Fatal(myerr)
	}
	fmt.Println(mymsg)

	Debug, LogLevel, startUpTime := getConfig()

	fmt.Println(Debug, LogLevel, startUpTime)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		// Create an error, covert the int type to a string
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}

func star(index int) (string, error) {
	if index < 0 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	stars := strings.Repeat("*", index)
	return stars, nil
}

func getConfig() (bool, string, time.Time) {
	return false, "INFO", time.Now()
}
