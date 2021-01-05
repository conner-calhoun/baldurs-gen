package main

import (
	"baldurs-gen/core"
	"fmt"
)

func main() {
	fmt.Println("Baldur's Gate 3 Generator")

	atlas := core.LoadAtlas("data/atlas.json")

	fmt.Println("Races: ", atlas.Races)
	fmt.Println("Classes: ", atlas.Classes[0].Name)
}
