package main

import (
	gg_bflow "github.com/alfarih31/gg-bflow/pkg/gg-bflow"
	"log"
)

func main() {
	err := gg_bflow.Init(".env")
	if err != nil {
		log.Fatal(err)
	}
	gg_bflow.Start()
}
