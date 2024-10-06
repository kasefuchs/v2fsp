package requester

type Config struct {
	URI    string `hcl:"uri"`    // Specifies destination uri.
	Method string `hcl:"method"` // Specifies http method to use.
}
