package apis

type Chart struct {
	Name    string `hcl:"name,label"`
	Version string `hcl:"version,attr"`
	Repo    Repo   `hcl:"repo,block"`
}
