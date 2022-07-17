package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configs struct {
	Db DB `yaml:"db"`
}

type DB struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslMode"`
}

func New() Configs {
	conf := Configs{}

	b, err := os.ReadFile("./configs.yaml")
	if err != nil {
		log.Fatalf("failed read configs.yaml. err :%s", err.Error())
	}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalf("failed Unmarshal configs.yaml. err :%s", err.Error())
	}

	return conf
}

func (c Configs) GetDb() DB {
	return c.Db
}
