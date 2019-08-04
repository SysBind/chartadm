package apis

type Chart struct {
	Name string `hcl:"name,label"`
	Repo Repo   `hcl:"repo,block"`
}
