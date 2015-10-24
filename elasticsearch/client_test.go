package elasticsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	assert := assert.New(t)

	cfg := NewConfig().WithURL("http://127.0.0.1:9200")
	client := NewClient(cfg)

	assert.NotNil(client)
}
