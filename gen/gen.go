package gen

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Character struct {
	Name     string
	Class    string
	Subclass string
	Party    []string
}

type Atlas struct {
	Races   []string `json:"races"`
	Classes []struct {
		Name string   `json:"name"`
		Subs []string `json:"subs"`
	} `json:"classes"`
	Party []string `json:"party"`
}

func LoadAtlas(atlasPath string) Atlas {
	var atlas Atlas

	jsonFile, _ := os.Open(atlasPath)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &atlas)

	return atlas
}

func RandomCharacter() Character {
	return Character{
		Name:     "Random",
		Class:    "Random",
		Subclass: "Random",
		Party:    []string{},
	}
}

func RandomRace(atlas Atlas) string {

	return ""
}
