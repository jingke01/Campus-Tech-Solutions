package main

import (
	"temp/config"
	"temp/database"
	"temp/global"
	"temp/router"
)

func main() {
	global.Conf = config.ConfigInit()
	global.DB = database.MysqlInit()

	router.RouterInit()
}
