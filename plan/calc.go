package plan

import (
	"fmt"
	"github.com/sysbind/chartadm/apis"
	"log"
)

// Calculate the plan according to desired and actual state
func Calc(cfg *apis.Config) (Plan, error) {
	log.Println("log.calculate: start")
	for i, bundle := range cfg.Bundles {
		fmt.Printf("Bundle %d: %s\n", i, bundle.Name)
		for j, release := range bundle.Releases {
			fmt.Printf("\tRelease %d: %s\n", j, release.Name)
			fmt.Printf("\t %s/%s\n", release.Chart.Repo.URL, release.Chart.Name)
		}
	}
	return Plan{}, nil
}
