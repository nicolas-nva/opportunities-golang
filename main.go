package main

import (
	"github.com/nicolas-nva/opportunities-go/config"
	"github.com/nicolas-nva/opportunities-go/router"
)

var(
	logger *config.Logger
)

func main() {
	logger =  config.GetLogger("main")
	// initialize configs
	err := config.Init()
	if err !=  nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	// initialize routes
	router.Initialize()
}