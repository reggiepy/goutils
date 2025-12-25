package confutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func TestJsonConfigBase(t *testing.T) {
	as := assert.New(t)
	base := JsonConfigBase{}
	cfg := TestConfig{Host: "localhost", Port: 8080}

	// Test ToJson
	jsonStr, err := base.ToJson(cfg)
	as.NoError(err)
	as.Contains(jsonStr, `"host": "localhost"`)
	as.Contains(jsonStr, `"port": 8080`)

	// Test LoadJson
	var loadedCfg TestConfig
	err = base.LoadJson(&loadedCfg, jsonStr)
	as.NoError(err)
	as.Equal(cfg, loadedCfg)

	// Test LoadJson Error
	err = base.LoadJson(&loadedCfg, "{invalid json")
	as.Error(err)
}

func TestYamlConfigBase(t *testing.T) {
	as := assert.New(t)
	base := YamlConfigBase{}
	cfg := TestConfig{Host: "127.0.0.1", Port: 9090}

	// Test ToYaml
	yamlStr, err := base.ToYaml(cfg)
	as.NoError(err)
	as.Contains(yamlStr, "host: 127.0.0.1")
	as.Contains(yamlStr, "port: 9090")

	// Test LoadYaml
	var loadedCfg TestConfig
	err = base.LoadYaml(&loadedCfg, yamlStr)
	as.NoError(err)
	as.Equal(cfg, loadedCfg)

	// Test LoadYaml Error
	err = base.LoadYaml(&loadedCfg, ": invalid yaml")
	as.Error(err)
}
