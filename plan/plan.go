package plan

import (
	//	"encoding/gob"
	"github.com/sysbind/chartadm/apis"
	//	"os"
)

type Operation int

const (
	install Operation = iota
	upgrade
	delete
)

type Step struct {
	release apis.Release
	op      Operation
}

type Plan struct {
	steps []Step
}

func (plan *Plan) Dump() {
	//enc := gob.NewEncoder(os.Stdout)
	//enc.Encode(plan)
}
