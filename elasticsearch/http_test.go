package elasticsearch

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	assert := assert.New(t)

	reader := strings.NewReader(`{"text":"Hello"}`)
	req := NewRequest("GET", "path", reader)
	assert.NotNil(req)
	assert.Equal("GET", req.method)
	assert.Equal("path", req.path)
	body, _ := ioutil.ReadAll(reader)
	data := map[string]interface{}{}
	json.Unmarshal(body, &data)
	assert.Equal("Hello", data["text"])
}
