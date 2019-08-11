package helm

import (
	"fmt"
	"github.com/sysbind/chartadm/apis"
)

type CLI struct {
	executable string
}

func init() {
	fmt.Println("helm.cli init..")
}

func (client *CLI) Fetch(chart apis.Chart) {
	fmt.Println("helm.cli fetch..")
}
