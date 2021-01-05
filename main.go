package main

import (
	"baldurs-gen/gen"
	"fmt"
)

func main() {
	fmt.Println("Baldur's Gate 3 Generator")

	atlas := gen.LoadAtlas("cfg/atlas.json")

	fmt.Println("Races: ", atlas.Races)
	fmt.Println("Classes: ", atlas.Classes[0].Name)
}
