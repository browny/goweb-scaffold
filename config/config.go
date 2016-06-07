// Package config does configuration management
package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// LoadAsset is wrapper function to read file from asset created by
// http://godoc.org/github.com/mjibson/esc
func LoadAsset(path string) http.File {
	asset := FS(false)
	file, _ := asset.Open(path)

	return file
}

// AppContext manages application context
type AppContext struct {
	Env          string
	ProjectID    string
	SuffixOfKind string
	EtcdServers  []string
	CommonConfig *viper.Viper
}

// Load loads config from config yaml file
func (ctx *AppContext) Load(env string) {
	log.Println("Load app context")

	// Load env specific config
	envConfig := viper.Sub(env)
	ctx.Env = env
	ctx.ProjectID = envConfig.GetString("project_id")
	ctx.SuffixOfKind = envConfig.GetString("datastore.kind_suffix")
	ctx.EtcdServers = envConfig.GetStringSlice("etcd")

	// Load common config
	ctx.CommonConfig = viper.Sub("common")
}

// Viper reads in viper config by scaning paths
func Viper() {
	// Load viper config
	viper.SetConfigName("viper")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("/go/src/goweb-scaffold/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	log.Printf("viper: %+v", viper.AllSettings())
}

// ViperHard reads in viper config by specified file
func ViperHard() {
	viper.SetConfigType("yaml")
	yaml, _ := ioutil.ReadAll(LoadAsset("/config/viper.yml"))

	err := viper.ReadConfig(bytes.NewBuffer(yaml))
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	log.Printf("viper: %+v", viper.AllSettings())
}
