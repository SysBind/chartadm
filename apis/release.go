package apis

// Represent Helm Release
type Release struct {
	Name  string `hcl:"name,label"`
	Chart Chart  `hcl:"chart,block"`
}
