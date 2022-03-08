package task

import (
	"bravo/config"
	"bravo/service"
)

func RunTasks(database service.Database, config config.Config) {
	currencyUpdater := NewCurrencyManager(database, config)
	currencyUpdater.updateCurrencyTask()
}
