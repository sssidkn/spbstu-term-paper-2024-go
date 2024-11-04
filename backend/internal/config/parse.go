package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Load(name string) *ConfigDB {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	cfg := &ConfigDB{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return cfg
}
