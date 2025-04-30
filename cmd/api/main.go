package main

import (
	"fmt"

	"github.com/Vagmacker/luzora-api/config"
	"github.com/Vagmacker/luzora-api/pkg/must"
)

func main() {
	cfg := must.Must(config.GetConfig())

	fmt.Println(cfg)
}
