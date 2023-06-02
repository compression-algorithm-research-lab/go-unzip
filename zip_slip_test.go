package unzip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsZipSlip(t *testing.T) {
	assert.False(t, IsZipSlip("./data", "foo"))
	assert.False(t, IsZipSlip("./data", "./foo"))
	assert.False(t, IsZipSlip("./data", "bar"))
	assert.False(t, IsZipSlip("./data", "bar/../foo"))
	assert.True(t, IsZipSlip("./data", "../foo"))
}
