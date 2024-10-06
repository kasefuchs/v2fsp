package requester

import (
	"github.com/gofiber/fiber/v2"
)

// Requester represents HTTP client dedicated to single internet resource.
type Requester struct {
	cfg *Config // Requester configuration.

	agt *fiber.Agent // Single use fiber.Agent of this Requester.
}

// New creates new instance of Requester from given Config.
func New(cfg *Config) *Requester {
	return &Requester{cfg: cfg}
}

// acquireAgent creates new fiber.Agent.
func (f *Requester) acquireAgent() error {
	f.agt = fiber.AcquireAgent()

	req := f.agt.Request()
	req.Header.SetMethod(f.cfg.Method)
	req.Header.SetRequestURI(f.cfg.URI)

	if err := f.agt.Parse(); err != nil {
		return err
	}

	return nil
}

// Fetch reads bytes from given internet resource.
func (f *Requester) Fetch() ([]byte, error) {
	if err := f.acquireAgent(); err != nil {
		return nil, err
	}

	_, body, errs := f.agt.Bytes()
	if len(errs) != 0 {
		return nil, errs[0]
	}

	return body, nil
}
