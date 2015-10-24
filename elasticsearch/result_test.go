package elasticsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	assert := assert.New(t)

	result := Result{}
	assert.NotNil(result)
}
