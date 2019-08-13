package plan

import (
	"fmt"
	"github.com/sysbind/chartadm/apis"
)

type Operation int

const (
	Install Operation = iota
	Upgrade
	Delete
)

type Step struct {
	Release apis.Release
	Op      Operation
}

func (op Operation) str() string {
	switch op {
	case Install:
		return "install"
	case Upgrade:
		return "upgrade"
	case Delete:
		return "delete"
	}
	return ""
}

type Plan struct {
	steps []Step
}

func (plan *Plan) Dump() {
	for _, step := range plan.steps {
		fmt.Printf("step: %s - %s\n", step.Release.Name, step.Op.str())
	}
}
