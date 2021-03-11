package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Character struct {
	Name     string   `json:"name"`
	Race     string   `json:"race"`
	Class    string   `json:"class"`
	Subclass string   `json:"subclass"`
	Party    []string `json:"party"`
}

type Atlas struct {
	Races   []string `json:"races"`
	Classes []struct {
		Name string   `json:"name"`
		Subs []string `json:"subs"`
	} `json:"classes"`
	Party []string `json:"party"`
}

func (c Character) Display() {
	fmt.Println("Name: ", c.Name)
	fmt.Println("Race: ", c.Race)
	fmt.Println("Class: ", c.Subclass, c.Class)
	fmt.Println("Party: ", c.Party)
}

func (a Atlas) RandomRace() string {
	return a.Races[0]
}

func (a Atlas) RandomClass() string {
	return ""
}

func (a Atlas) RandomSubclass() string {
	// NOTE: This depends on a class already being chosen
	return ""
}

func (a Atlas) RandomParty() []string {
	return []string{}
}

func LoadAtlas(atlasPath string) Atlas {
	var atlas Atlas

	jsonFile, _ := os.Open(atlasPath)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &atlas)

	return atlas
}
