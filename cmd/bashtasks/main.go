package main

import (
	"log"

	"github.com/saromanov/bashtasks/bashtasks"
)

func main() {
	cfg, err := bashtasks.LoadYAML("../../configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if len(cfg.Tasks) > 0 {
		bashtasks.ExecuteRowTasks(cfg.Tasks)
	}
}
