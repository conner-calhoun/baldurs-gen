package main

import (
	"baldurs-gen/core"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Baldur's Gate 3 Generator")

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// TODO: verify that all files exist in the data folder, possibly combine json files
	// TODO: combine all config/data files into a single atlas
	myChar := core.RandomCharacter("./data")
	myChar.Display()
}
