package task

import (
	"bravo/config"
	"bravo/repository"
)

func RunTasks(database repository.Database, config config.Config) {
	currencyUpdater := NewCurrencyManager(database, config)
	currencyUpdater.updateCurrencyTask()
}
