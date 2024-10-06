package updater

import "github.com/kasefuchs/v2fsp/internal/pkg/requester"

type Config struct {
	Cron string `hcl:"cron"`

	Request requester.Config `hcl:"request,block"`
}
