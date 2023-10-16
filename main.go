package main

import (
	"github.com/Guilherme-daCosta/gopportunities/config"
	"github.com/Guilherme-daCosta/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errf("config initilization erro: %v", err)
		return
	}

	router.Initialize()
}
