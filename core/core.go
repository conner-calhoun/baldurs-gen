package core

func RandomCharacter() Character {
	// Read all atlases

	return Character{
		Name:     "Random",   // NameAtlas.Random()
		Gender:   "M|F|T",    // RandomGender()
		Race:     "Random",   // Atlas.RandomRace()
		Class:    "Random",   // Atlas.RandomClass()
		Subclass: "Random",   // Atlas.RandomSubclass()
		Party:    []string{}, // Atlas.RandomParty()
	}
}
