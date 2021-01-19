package creature

import (
	"math/rand"
	"time"
)

var creatures = []string{"shark", "jellyfish", "squid", "octopus", "dolphin"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random function exported outside
func Random() string {
	i := rand.Intn(len(creatures))
	return creatures[i]
}
