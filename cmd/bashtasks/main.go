package main

import (
	"log"
	"os"

	"github.com/saromanov/bashtasks/bashtasks"
)

// getFileName return file name in the case
// if this is defined as a command line argument
// for example
// > bashtasks foo
// It'll load file foo.yaml from the BASHTASKS env.path
func getFileName() string {
	args := os.Args
	if len(args) <= 1 {
		return ""
	}
	return args[1]
}

func main() {
	name := getFileName()
	cfg := &bashtasks.Config{}
	var err error
	if name != "" {
		cfg, err = bashtasks.LoadYAMLByName(name)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		cfg, err = bashtasks.LoadYAML("../../configs/config.yaml")
		if err != nil {
			log.Fatal(err)
		}
	}
	bt := bashtasks.New(cfg)

	bashtasks.StartMessage(cfg)
	if len(cfg.Tasks) > 0 {
		bt.Run()
	}
	bt.Response()
}
