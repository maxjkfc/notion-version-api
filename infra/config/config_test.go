package config

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func Test_ReadConfig(t *testing.T) {
	c, err := New("test.yaml")
	assert.NoError(t, err)
	spew.Dump(c)
}
