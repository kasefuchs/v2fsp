package v2fsp

import (
	"github.com/kasefuchs/v2fsp/internal/pkg/server"
	"github.com/kasefuchs/v2fsp/internal/pkg/updater"
)

// Config represents V2FSP configuration.
type Config struct {
	Server server.Config    `hcl:"server,block"` // Server listener configuration.
	Source []updater.Config `hcl:"source,block"` // Source specifies sources from which to get outbound links.
}
