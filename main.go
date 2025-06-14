package main

import (
	"fmt"

	"github.com/nicolas-nva/opportunities-go/config"
	"github.com/nicolas-nva/opportunities-go/router"
)

func main() {
	
	// initialize configs
	err := config.Init()
	if err !=  nil {
		fmt.Println(err)
		return
	}

	// initialize routes
	router.Initialize()
}