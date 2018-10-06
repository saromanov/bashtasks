package bashtasks

import (
	"fmt"

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

// LoadYAML provides loading of the yaml config
func LoadYAML(path string) error {
	cfg := &Config{}
	_, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	fmt.Println(cfg)
	return nil
}
