package system

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Configuration struct {
	sqliteConf string `yaml:"db"` //database
}

const (
	DEFAULT_PAGESIZE = 10
)

var configuration *Configuration

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	/*if config.PageSize <= 0 {
		config.PageSize = DEFAULT_PAGESIZE
	}*/
	configuration = &config
	return err
}

func GetConfiguration() *Configuration {
	return configuration
}
