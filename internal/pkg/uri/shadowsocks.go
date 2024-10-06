package uri

import (
	"bytes"
	"encoding/base64"
	"errors"
	"net/url"
	"strings"

	"github.com/v2fly/v2ray-core/v5/app/subscription/specs"
	"github.com/v2fly/v2ray-core/v5/infra/conf/jsonpb"
	"github.com/v2fly/v2ray-core/v5/proxy/shadowsocks"
	"github.com/v2fly/v2ray-core/v5/proxy/shadowsocks/simplified"
)

// extractShadowsocksAccount extracts shadowsocks.Account from url.URL.
func extractShadowsocksAccount(uri *url.URL) (*shadowsocks.Account, error) {
	uid, err := base64.URLEncoding.DecodeString(uri.User.Username())
	if err != nil {
		return nil, err
	}

	uip := strings.SplitN(string(uid), ":", 2)
	if len(uip) != 2 {
		return nil, errors.New("incorrect amount of user info segments")
	}

	cipher := shadowsocks.CipherFromString(uip[0])
	if cipher == shadowsocks.CipherType_UNKNOWN {
		return nil, errors.New("unknown cipher type")
	}

	return &shadowsocks.Account{
		CipherType: cipher,
		Password:   uip[1],
	}, nil
}

// ParseShadowsocks parses specs.OutboundConfig from url.URL
func ParseShadowsocks(uri *url.URL) (*specs.OutboundConfig, error) {
	ac, err := extractShadowsocksAccount(uri)
	if err != nil {
		return nil, err
	}

	endpoint, err := parseServerEndpoint(uri, nil)
	if err != nil {
		return nil, err
	}

	config := &simplified.ClientConfig{
		Address:  endpoint.Address,
		Port:     endpoint.Port,
		Password: ac.Password,
		Method:   &simplified.CipherTypeWrapper{Value: ac.CipherType},
	}

	buf := bytes.NewBuffer(nil)
	if err := jsonpb.DumpJSONPb(config, buf); err != nil {
		return nil, err
	}

	return &specs.OutboundConfig{
		Protocol: "shadowsocks",
		Settings: buf.Bytes(),
	}, nil
}
