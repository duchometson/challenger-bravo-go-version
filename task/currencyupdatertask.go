package task

import (
	"bravo/service"
	"fmt"
	"strings"
	"time"

	exchange "github.com/asvvvad/exchange"
)

type CurrencyUpdater struct {
	database    service.Database
	finishTasks chan bool
}

var mainServerCurrency string
var amount int

func (c *CurrencyUpdater) updateCurrencyTask() {
	fmt.Println("CurrencyUpdater - Started UpdateTask")
	ticker := time.NewTicker(5 * time.Second)
	mainServerCurrency = "USD"
	amount = 1
	for {
		<-ticker.C
		currenciesList := c.database.GetAllCurrencies()
		c.updateAllCurrencies(currenciesList)
	}
}

func (c *CurrencyUpdater) updateAllCurrencies(currenciesList []string) {
	for _, currency := range currenciesList {
		strings.ReplaceAll(currency, " ", "")
		currencyToUpdate := exchange.New(currency)
		value, err := currencyToUpdate.ConvertTo(mainServerCurrency, amount)
		if err == nil {
			valueAsFloat, _ := value.Float64()
			c.database.InsertOrUpdate(currency, valueAsFloat)
		} else {
			fmt.Println(err)
		}
	}
}

func NewCurrencyManager(database service.Database) *CurrencyUpdater {
	return &CurrencyUpdater{
		database: database,
	}
}
