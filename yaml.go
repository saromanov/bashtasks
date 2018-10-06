package main

import (
	"fmt"
	"log"

	"github.com/go-yaml/yaml"
)

// Config defines struct for config
type Config struct {
	Tasks []Task `yaml:"tasks"`
}

// Task is definition for task
type Task struct {
	Title string `yaml:"title"`
	Cmd   string `yaml:"cmd"`
}

func loadYAML(path string) {
	cfg := &Config{}
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(d)
}
