package repository

import (
	"github.com/goccy/go-json"
	"github.com/v2fly/v2ray-core/v5/app/subscription/specs"
	"golang.org/x/crypto/sha3"
)

// Repository contains specs.OutboundConfig.
type Repository map[[32]byte]*specs.OutboundConfig

// Add pushes single config entry to Repository.
func (r *Repository) Add(config *specs.OutboundConfig) error {
	marshal, err := json.Marshal(config)
	if err != nil {
		return err
	}

	hash := sha3.Sum256(marshal)
	(*r)[hash] = config

	return nil
}

// Items returns all items inside repository.
func (r *Repository) Items() []*specs.OutboundConfig {
	index := 0
	items := make([]*specs.OutboundConfig, len(*r))

	for _, item := range *r {
		items[index] = item
		index++
	}

	return items
}
