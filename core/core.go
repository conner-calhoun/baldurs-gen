package core

import (
	"fmt"
	"math/rand"
	"time"
)

// Return a random item from a list of strings
func GetRandom(items []string) string {
	rand.Seed(time.Now().Unix())
	return items[rand.Intn(len(items))]
}

func RandomCharacter(dataPath string) Character {
	// Read all atlases
	baseAtlas := LoadAtlas(fmt.Sprintf("%s/atlas.json", dataPath))
	nameAtlas := LoadNameAtlas(fmt.Sprintf("%s/names.json", dataPath))

	return Character{
		Name:     nameAtlas.RandomName(),     // NameAtlas.Random()
		Gender:   RandomGender(),             // RandomGender()
		Race:     baseAtlas.RandomRace(),     // Atlas.RandomRace()
		Class:    baseAtlas.RandomClass(),    // Atlas.RandomClass()
		Subclass: baseAtlas.RandomSubclass(), // Atlas.RandomSubclass()
		Party:    baseAtlas.RandomParty(),    // Atlas.RandomParty()
	}
}
