package gen

import (
	"fmt"
)

type NameAtlas struct {
	FirstHalf  []string
	LastHalf   []string
	Conjuct    []string
	Descriptor []string
}

type Name struct {
	FirstHalf  string
	LastHalf   string
	Conjuct    string
	Descriptor string
}

func (n Name) Build() string {
	return fmt.Sprintf("%s%s %s %s",
		n.FirstHalf,
		n.LastHalf,
		n.Conjuct,
		n.Descriptor)
}
