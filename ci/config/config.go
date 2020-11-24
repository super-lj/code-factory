package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type RepoConfig struct {
	Path string
	Img  string
}

type ImageConfig struct {
	Path string
	Tag  string
}

type Config struct {
	Repos  map[string]RepoConfig
	Images map[string]ImageConfig
}

var CurrConfig Config

func InitCIConfig(configFilePath string) error {
	// read the yml config file
	bstr, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(bstr), &CurrConfig)
	if err != nil {
		return err
	}

	// fix all relative paths
	for name, config := range CurrConfig.Repos {
		if filepath.IsAbs(config.Path) {
			continue
		}
		CurrConfig.Repos[name] = RepoConfig{
			Path: filepath.Join(filepath.Dir(configFilePath), config.Path),
			Img:  config.Img,
		}
	}
	for name, config := range CurrConfig.Images {
		if filepath.IsAbs(config.Path) {
			continue
		}
		CurrConfig.Images[name] = ImageConfig{
			Path: filepath.Join(filepath.Dir(configFilePath), config.Path),
			Tag:  config.Tag,
		}
	}
	return nil
}
