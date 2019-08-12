package helm

import (
	"fmt"
	"github.com/sysbind/chartadm/apis"
)

type CLI struct {
	Config apis.ChartAdmConfig
}

func Init(executable string) {
	fmt.Println("helm.cli init..")

}

func (client *CLI) Fetch(chart apis.Chart) {
	fmt.Println("helm.cli fetch..")
}
