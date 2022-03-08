package main

import (
	"bravo/config"
	"bravo/controller"
	"bravo/dao"
	"bravo/task"
)

func main() {
	database := dao.NewMockedCoins()
	config := config.NewConfigutaror()
	go controller.InitializeServerRoutes(database)
	task.RunTasks(database, config)
}
