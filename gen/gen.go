package gen

func RandomCharacter() Character {
	// Read all atlases

	return Character{
		Name:     "Random",   // NameAtlas.Random()
		Race:     "Random",   // Atlas.RandomRace()
		Class:    "Random",   // Atlas.RandomClass()
		Subclass: "Random",   // Atlas.RandomSubclass()
		Party:    []string{}, // Atlas.RandomParty()
	}
}
