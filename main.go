package main

import (
	"github.com/mariojuzar/soldier-magazine/api"
	"github.com/mariojuzar/soldier-magazine/api/configuration"
)

func main() {
	engine := api.Run()
	_ = engine.Run(":" + configuration.Port)
}
