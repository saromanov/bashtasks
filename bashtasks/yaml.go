package bashtasks

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config defines struct for config
type Config struct {
	Tasks         []Task `yaml:"tasks"`
	ParallelTasks []Task `yaml:"parallel_tasks"`
}

// Task is definition for task
type Task struct {
	Title string `yaml:"title"`
	Cmd   string `yaml:"cmd"`
	// AbortPipeline stops whole pipeline
	// in the case if was occured error
	AbortPipeline bool `yaml:"abort_pipeline"`
}

// LoadYAML provides loading of the yaml config
func LoadYAML(path string) (*Config, error) {

	cfg := &Config{}
	fileConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(fileConfig, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
