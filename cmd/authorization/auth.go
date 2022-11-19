package main

import (
	mcsauth "eventool/internal/app/authorization"
	"eventool/internal/pkg/utils/config"
	"eventool/internal/pkg/utils/log"
)

func main() {
	
	err := config.DevConfigStore.FromJson()
	if err != nil {
		log.Error(err)
	}
	
	err = config.ProdConfigStore.FromJson()
	if err != nil {
		log.Error(err)
	}

	mcsauth.RunServer()
}
