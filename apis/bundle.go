/*
Copyright © 5769 (2019) Asaf Ohayon <asaf@sysbind.co.il>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apis

import (
	"fmt"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"
	"os"
)

// Represent Bundle of Helm Release
type Bundle struct {
	Name     string    `hcl:"name,label"`
	Releases []Release `hcl:"release,block"`
}

type Config struct {
	Bundles []Bundle `hcl:"bundle,block"`
}

// Type definitions from above ommitted for brevity

func ParseConfig() (*Config, error) {
	c := Config{}

	parser := hclparse.NewParser()

	f, diags := parser.ParseHCLFile("docs/sample.cfg")

	if diags.HasErrors() {
		fmt.Println("diags has errors (parse)")
		return nil, fmt.Errorf("DecodeBody failed", f)
	}
	wr := hcl.NewDiagnosticTextWriter(os.Stdout, parser.Files(), 78, true)
	wr.WriteDiagnostics(diags)
	diags = gohcl.DecodeBody(f.Body, nil, &c)
	if diags.HasErrors() {
		fmt.Println("diags has errors (decode)")
		return nil, fmt.Errorf("DecodeBody failed", f)
	}
	wr.WriteDiagnostics(diags)

	for i, bundle := range c.Bundles {
		fmt.Printf("Bundle %d: %s\n", i, bundle.Name)
		for j, release := range bundle.Releases {
			fmt.Printf("\tRelease %d: %s\n", j, release.Name)
			fmt.Printf("\t %s/%s\n", release.Chart.Repo.URL, release.Chart.Name)
		}
	}

	return &c, nil
}
