package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var fileCnf *Configurations = nil

func GetConfigFromYaml() (*Configurations, error) {
	if fileCnf == nil {
		yamlFile, err := ioutil.ReadFile("config/configs.yaml")
		if err != nil {
			return nil, err
		}
		temp := &Configurations{}
		err = yaml.Unmarshal(yamlFile, temp)
		if err != nil {
			return nil, err
		}
		fileCnf = temp

	}
	return fileCnf, nil
}

type Configurations struct {
	AuthenticationConfiguration AuthenticationConfiguration `yaml:"AuthenticationConfiguration"`
	PanelServerConfiguration    PanelServerConfiguration    `yaml:"PanelServer"`
}

type AuthenticationConfiguration struct {
	TokenSecret string `yaml:"tokenSecret"`
}

type PanelServerConfiguration struct {
	Host     string `yaml:"host"`
	HTTPPort string `yaml:"httpPort"`
}
