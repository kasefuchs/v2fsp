package server

// Config represents HTTP-server configuration.
type Config struct {
	Address string `hcl:"address"` // The address to bind the srv.
	Port    uint16 `hcl:"port"`    // The port to bind the srv.
}
