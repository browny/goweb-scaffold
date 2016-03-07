// Package config does configuration management
package config

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GlobalConfig is the binding object of config json file
type GlobalConfig struct {
	ProjectId string `json:"project_id"`
}

// LoadGcloudConfig loads config.json to GlobalConfig
func LoadGcloudConfig(file http.File) GlobalConfig {
	decoder := json.NewDecoder(file)
	configuration := GlobalConfig{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}

// LoadAsset is wrapper function to read file from asset created by
// http://godoc.org/github.com/mjibson/esc
func LoadAsset(path string) http.File {
	asset := FS(false)
	file, _ := asset.Open(path)

	return file
}
