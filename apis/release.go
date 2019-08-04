package apis

// Represent Helm Release
type Release struct {
	Name  string `hcl:"name,label"`
	Chart string `hcl:"chart,attr"`
}

// Type definitions from above ommitted for brevity

/*func ParseRelease(hclText string) (*Release, error) {
	result := &Release{}

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
*/
