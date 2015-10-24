package elasticsearch

import (
	"net/http"
	"net/url"
)

type Config struct {
	url string

	// The maximum number of times that a request will be retried for failures.
	// Defaults to -1, which defers the max retry setting to the service specific
	// configuration.
	MaxRetries int

	// Healthchecks enabled or disabled
	HealthcheckEnabled bool

	// Sniffer enabled or disabled
	SnifferEnabled bool

	// The HTTP client to use when sending requests.
	HTTPClient *http.Client
}

// NewConfig returns a new Config pointer
func NewConfig() *Config {
	return &Config{
		HealthcheckEnabled: true,
		SnifferEnabled:     true,
		HTTPClient:         http.DefaultClient,
	}
}

func (c *Config) WithURL(url string) *Config {
	c.url = url
	return c
}

func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}

func (c *Config) GetURL() string {
	return c.url
}

func (c *Config) GetScheme() string {
	url, err := url.Parse(c.url)
	if err != nil {
		return ""
	}
	return url.Scheme
}

func (c *Config) IsSSL() bool {
	return c.GetScheme() == "https"
}
