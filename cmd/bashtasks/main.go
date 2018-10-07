package main

import (
	"fmt"
	"log"

	"github.com/saromanov/bashtasks/bashtasks"
)

func main() {
	cfg, err := bashtasks.LoadYAML("../../configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
