package uri

import (
	"errors"
	"net/url"

	"github.com/v2fly/v2ray-core/v5/app/subscription/specs"
)

// Parse parses specs.OutboundConfig from url.URL.
func Parse(uri *url.URL) (*specs.OutboundConfig, error) {
	switch uri.Scheme {
	case "ss":
		return ParseShadowsocks(uri)
	default:
		return nil, errors.New("unknown protocol scheme: " + uri.Scheme)
	}
}
