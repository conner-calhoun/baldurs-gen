package main

import (
	"baldurs-gen/core"
	"fmt"
)

func main() {
	fmt.Println("Baldur's Gate 3 Generator")
	myChar := core.RandomCharacter()
	myChar.Display()
}
