package helm

import "github.com/sysbind/chartadm/apis"

type Client interface {
	Fetch(chart apis.Chart) error
}
