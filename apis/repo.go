package apis

// Represent Helm Repo
type Repo struct {
	Name string `hcl:"name,label"`
	URL  string `hcl:"url,attr"`
}
