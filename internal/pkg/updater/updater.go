package updater

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"github.com/go-co-op/gocron/v2"
	"github.com/kasefuchs/v2fsp/internal/pkg/repository"
	"github.com/kasefuchs/v2fsp/internal/pkg/requester"
	"github.com/kasefuchs/v2fsp/internal/pkg/uri"
	"net/url"
)

type Updater struct {
	cfg *Config // Updater configuration.

	req *requester.Requester // Requester of this Updater.
}

func New(cfg *Config) *Updater {
	req := requester.New(&cfg.Request)

	return &Updater{cfg: cfg, req: req}
}

func (u *Updater) Schedule(scheduler gocron.Scheduler) (gocron.Job, error) {
	return scheduler.NewJob(
		gocron.CronJob(u.cfg.Cron, false),
		gocron.NewTask(u.RunTask),
	)
}

func (u *Updater) RunTask() error {
	fetch, err := u.req.Fetch()
	if err != nil {
		return err
	}

	decodeString, err := base64.StdEncoding.DecodeString(string(fetch))
	if err != nil {
		return err
	}

	reader := bytes.NewReader(decodeString)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		txt := scanner.Text()

		u, err := url.Parse(txt)
		if err != nil {
			continue
		}

		oc, err := uri.Parse(u)
		if err != nil {
			continue
		}

		if err := repository.Add(oc); err != nil {
			continue
		}
	}

	return nil
}
