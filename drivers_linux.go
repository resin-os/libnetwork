package libnetwork

import (
	"github.com/docker/libnetwork/drivers/bridge"
	"github.com/docker/libnetwork/drivers/host"
	"github.com/docker/libnetwork/drivers/null"
)

func getInitializers(experimental bool) []initializer {
	in := []initializer{
		{bridge.Init, "bridge"},
		{host.Init, "host"},
		{null.Init, "null"},
	}
	return in
}
