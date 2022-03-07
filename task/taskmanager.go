package task

import "bravo/service"

func RunTasks(database service.Database) {
	currencyUpdater := NewCurrencyManager(database)
	currencyUpdater.updateCurrencyTask()
}
