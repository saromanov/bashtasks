package bashtasks

import (
	"io/ioutil"

	"github.com/saromanov/cowrow"

	"gopkg.in/yaml.v2"
)

const envPath = "BASHTASKS"

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

	// Tags provides marking of the task
	Tags []string `yaml:"tags"`
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

// LoadYAMLByName provides loading of the yaml file
// from directory by the name
func LoadYAMLByName(name string) (*Config, error) {
	cfg := &Config{}
	err := cowrow.LoadYAMLFile(envPath, name, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
