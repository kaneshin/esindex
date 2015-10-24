package elasticsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := NewConfig()
	assert.NotNil(cfg)

	var url = "http://example.com"

	cfg.WithURL(url)
	assert.False(cfg.IsSSL())
	assert.Equal(url, cfg.GetURL())

	url = "https://example.com"
	cfg.WithURL(url)
	assert.True(cfg.IsSSL())
	assert.Equal(url, cfg.GetURL())
}
