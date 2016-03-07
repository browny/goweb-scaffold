package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadGcloudConfig(t *testing.T) {
	config := LoadGcloudConfig(LoadAsset("/config/config.json"))

	assert.Equal(t, "goweb-scaffold", config.ProjectId)
}
