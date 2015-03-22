package config

// Config contains RPC configuration for a Namecoin node.
type Config struct {
	User     string
	Password string
	Server   string
}

// New returns a new Config
func New(user, pass, server string) Config {
	return Config{
		User:     user,
		Password: pass,
		Server:   server,
	}
}
