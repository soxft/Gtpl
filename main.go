package main

import (
	"gtpl/config"
	"gtpl/library/zag"
	"gtpl/process/api"
	"gtpl/process/db"
	"gtpl/process/rdb"
)

func main() {
	config.Init()
	zag.Init()

	// 初始化 redis 和 mysql
	rdb.Init()
	db.Init()

	api.Run()
}
