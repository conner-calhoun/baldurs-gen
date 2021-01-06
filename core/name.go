package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type NameAtlas struct {
	FirstHalf  []string
	LastHalf   []string
	Descriptor []string
}

type Name struct {
	FirstHalf  string
	LastHalf   string
	Descriptor string // Could either be from the list provided in the data... or "of ${city}"
}

func (n Name) Build() string {
	return fmt.Sprintf("%s%s %s",
		n.FirstHalf,
		n.LastHalf,
		n.Descriptor)
}

func (n NameAtlas) RandomName() string {
	return Name{}.Build()
}

func LoadNameAtlas(atlasPath string) NameAtlas {
	var atlas NameAtlas

	jsonFile, _ := os.Open(atlasPath)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &atlas)

	return atlas
}
