package main

import (
	"baldurs-gen/core"
	"fmt"
)

func main() {
	fmt.Println("Baldur's Gate 3 Generator")

	// TODO: verify that all files exist in the data folder, possibly combine json files

	myChar := core.RandomCharacter("data")
	myChar.Display()
}
