package apis

import (
	"fmt"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"
	"os"
)

// Represent Helm Release
type Bundle struct {
	Name     string    `hcl:"name,label"`
	Releases []Release `hcl:"release,block"`
}

type Config struct {
	Bundles []Bundle `hcl:"bundle,block"`
}

// Type definitions from above ommitted for brevity

func ParseBundle(hclText string) (*Bundle, error) {
	c := Config{}

	parser := hclparse.NewParser()

	f, diags := parser.ParseHCLFile("docs/sample.cfg")

	if diags.HasErrors() {
		fmt.Println("diags has errors (parse)")
	}
	wr := hcl.NewDiagnosticTextWriter(os.Stdout, parser.Files(), 78, true)
	wr.WriteDiagnostics(diags)
	diags = gohcl.DecodeBody(f.Body, nil, &c)
	if diags.HasErrors() {
		fmt.Println("diags has errors (decode)")
	}
	wr.WriteDiagnostics(diags)

	for i, bundle := range c.Bundles {
		fmt.Printf("Bundle %d: %s\n", i, bundle.Name)
		for j, release := range bundle.Releases {
			fmt.Printf("\tRelease %d: %s\n", j, release.Name)
			fmt.Printf("\t %s/%s\n", release.Chart.Repo.URL, release.Chart.Name)
		}
	}

	return nil, nil
}
