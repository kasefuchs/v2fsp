package uri

import (
	"net/url"
	"strconv"

	"github.com/v2fly/v2ray-core/v5/common/net"
	"github.com/v2fly/v2ray-core/v5/common/protocol"
)

// parseIPOrDomain parses net.IPOrDomain from url.URL.
func parseIPOrDomain(uri *url.URL) *net.IPOrDomain {
	hostname := uri.Hostname()
	ip := net.ParseIP(hostname)

	iord := &net.IPOrDomain{}
	if ip != nil {
		iord.Address = &net.IPOrDomain_Ip{Ip: ip}
	} else {
		iord.Address = &net.IPOrDomain_Domain{Domain: hostname}
	}

	return iord
}

// parseServerEndpoint parses protocol.ServerEndpoint from url.URL.
func parseServerEndpoint(uri *url.URL, user []*protocol.User) (*protocol.ServerEndpoint, error) {
	port, err := strconv.Atoi(uri.Port())
	if err != nil {
		return nil, err
	}

	return &protocol.ServerEndpoint{
		Address: parseIPOrDomain(uri),
		Port:    uint32(port),
		User:    user,
	}, nil
}
