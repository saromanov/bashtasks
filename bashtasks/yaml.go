package bashtasks

import (
	"fmt"

	"gopkg.in/yaml.v2"
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
	cfg := Config{}
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	fmt.Println(string(d))
	fmt.Println(cfg)
	return nil
}
