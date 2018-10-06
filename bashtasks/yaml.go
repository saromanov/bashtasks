package bashtasks

import (
	"fmt"
	"io/ioutil"

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
	fileConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileConfig, &cfg)
	if err != nil {
		return err
	}
	fmt.Println(cfg)
	return nil
}
