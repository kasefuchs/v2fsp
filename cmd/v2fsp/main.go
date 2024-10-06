package main

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/kasefuchs/v2fsp/internal/app/v2fsp"
	"github.com/kasefuchs/v2fsp/internal/app/v2fsp/route"
	"github.com/kasefuchs/v2fsp/internal/pkg/server"
	"github.com/kasefuchs/v2fsp/internal/pkg/updater"
)

func main() {
	cfg := &v2fsp.Config{}
	if err := hclsimple.DecodeFile("config.hcl", nil, cfg); err != nil {
		panic(err)
	}

	she, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	for _, sc := range cfg.Source {
		upd := updater.New(&sc)
		_ = upd.RunTask()
		if _, err := upd.Schedule(she); err != nil {
			panic(err)
		}
	}

	she.Start()

	server.Use(route.SubscriptionRoute)
	if err := server.Listen(&cfg.Server); err != nil {
		panic(err)
	}
}
