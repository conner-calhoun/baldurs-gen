package core

// Return a random item from a list of strings
func GetRandom(items []string) string {

	return ""
}

func RandomCharacter() Character {
	// Read all atlases
	baseAtlas := LoadAtlas("data/atlas.json")
	nameAtlas := LoadNameAtlas("data/names.json")

	return Character{
		Name:     nameAtlas.RandomName(),     // NameAtlas.Random()
		Gender:   RandomGender(),             // RandomGender()
		Race:     baseAtlas.RandomRace(),     // Atlas.RandomRace()
		Class:    baseAtlas.RandomClass(),    // Atlas.RandomClass()
		Subclass: baseAtlas.RandomSubclass(), // Atlas.RandomSubclass()
		Party:    baseAtlas.RandomParty(),    // Atlas.RandomParty()
	}
}
