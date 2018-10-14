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
	bt := bashtasks.New(cfg)

	bashtasks.StartMessage(cfg)
	if len(cfg.Tasks) > 0 {
		bt.ExecuteRowTasks()
	}
}
