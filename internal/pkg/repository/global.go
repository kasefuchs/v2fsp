package repository

import (
	"github.com/v2fly/v2ray-core/v5/app/subscription/specs"
)

// Global repository.
var globalRepository = &Repository{}

// Add pushes given configs to repository.
func Add(configs ...*specs.OutboundConfig) []error {
	var errors []error
	for _, config := range configs {
		err := globalRepository.Add(config)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

// Items returns items in globalRepository repository.
func Items() []*specs.OutboundConfig {
	return globalRepository.Items()
}

// Length returns amount of items in globalRepository repository.
func Length() int {
	return len(*globalRepository)
}
