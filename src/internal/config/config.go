package config

const Version = "0.0.2"

type Config struct {
	Debug       bool
	HTTPPort    uint64
	StaticPath  string
	UseSSL      bool
	SSLCertPath string
	SSLKeyPath  string
}
