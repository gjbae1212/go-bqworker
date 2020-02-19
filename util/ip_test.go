package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetHostname(t *testing.T) {
	assert := assert.New(t)
	hostname, err := os.Hostname()
	assert.NoError(err)
	tests := map[string]struct {
		output string
	}{
		"success": {output: hostname},
	}

	for _, t := range tests {
		assert.Equal(t.output, GetHostname())
	}
}
