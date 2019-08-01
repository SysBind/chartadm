package apis

import (
	"github.com/hashicorp/hcl"
	"log"
)

// Represent Helm Release
type Bundle struct {
	Name    string    `hcl:"name,label"`
	Release []Release `hcl:"release,block`
}

// Type definitions from above ommitted for brevity

func ParseBundle(hclText string) (*Bundle, error) {
	result := &Bundle{}

	hclParseTree, err := hcl.Parse(hclText)
	if err != nil {
		return nil, err
	}

	if rawName := hclParseTree.Get("release", false); rawName != nil {
		result.Name = rawName.Value.(string)
	}

	log.Printf("%+v\n", result)
	return nil, nil
}
