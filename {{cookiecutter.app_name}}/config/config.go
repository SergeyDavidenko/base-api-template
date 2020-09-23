package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

// Conf ...
var Conf ConfYaml
var defaultConf = []byte(`
core:
  environment: "dev"
  lable: "{{cookiecutter.app_name}}"
api:
  port: ":8088"
  health_port: ":1499"
  metric_uri: "/metrics"
  health_uri: "/healthz"
log:
  level: "info"
  api_level: "warning"
`)

// ConfYaml is config structure.
type ConfYaml struct {
	Core          SectionCore          `yaml:"core"`
	API           SectionAPI           `yaml:"api"`
	Log           SectionLog           `yaml:"log"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Environment string `yaml:"environment"`
	Lable       string `yaml:"lable"`
}

// SectionAPI is sub section of config.
type SectionAPI struct {
	MetricURI  string `yaml:"metric_uri"`
	HealthURI  string `yaml:"health_uri"`
	Port       string `yaml:"port"`
	HealthPort string `yaml:"health_port"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Level    string `yaml:"level"`
	APILevel string `yaml:"api_level"`
}


// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml
	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("go")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory with name ".gorush" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}

	// Core
	conf.Core.Environment = viper.GetString("core.environment")
	conf.Core.Lable = viper.GetString("core.lable")

	//API
	conf.API.Port = viper.GetString("api.port")
	conf.API.HealthPort = viper.GetString("api.health_port")
	conf.API.HealthURI = viper.GetString("api.health_uri")
	conf.API.MetricURI = viper.GetString("api.metric_uri")

	//Log
	conf.Log.Level = viper.GetString("log.level")
	conf.Log.APILevel = viper.GetString("log.api_level")


	return conf, nil
}