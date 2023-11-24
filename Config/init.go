package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

func init() {
	cfg_path := path.Join("Config", "config.yaml")
	// if no config.yaml, create a config.yaml
	if _, err := os.Stat(cfg_path); os.IsNotExist(err) {
		f, err := os.Create(cfg_path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		buf, err := yaml.Marshal(Cfg)
		if err != nil {
			panic(err)
		}

		_, err = f.Write(buf)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.Open(cfg_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(Cfg)
	if err != nil {
		panic(err)
	}
}
