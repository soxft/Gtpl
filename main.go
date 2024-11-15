package main

import (
	"gtpl/config"
	"gtpl/library/zag"
	"gtpl/process/api"
)

func main() {
	zag.Init()
	config.Init()

	api.Run()
}
