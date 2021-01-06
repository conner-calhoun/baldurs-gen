package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Character struct {
	Name     string
	Gender   string
	Race     string
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

func (c Character) Display() {
	fmt.Println("Name: ", c.Name)
	fmt.Println("Race: ", c.Race)
	fmt.Println("Class: ", c.Subclass, c.Class)
	fmt.Println("Party: ", c.Party)
}

func RandomRace(atlas Atlas) string {
	return ""
}

func LoadAtlas(atlasPath string) Atlas {
	var atlas Atlas

	jsonFile, _ := os.Open(atlasPath)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &atlas)

	return atlas
}
