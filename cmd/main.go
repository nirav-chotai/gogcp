package main

import (
	"fmt"
	"time"

	"github.com/nirav-chotai/gogcp/creature"
	"github.com/nirav-chotai/gogcp/logging"
)

func main() {

	logger := logging.New(time.RFC3339, true)

	logger.Log("info", "starting up service")
	logger.Log("warning", "no tasks found")
	logger.Log("error", "exiting: no work performed")

	fmt.Println(creature.Random())
	fmt.Println(creature.Random())
	fmt.Println(creature.Random())
	fmt.Println(creature.Random())
}
