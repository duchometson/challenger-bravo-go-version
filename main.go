package main

import (
	"bravo/controller"
	"bravo/dao"
	"bravo/task"
)

func main() {
	database := dao.NewMockedCoins()
	go controller.InitializeServerRoutes(database)
	task.RunTasks(database)
}
