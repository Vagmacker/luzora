package main

import (
	"fmt"

	"github.com/Vagmacker/luzora-api/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
