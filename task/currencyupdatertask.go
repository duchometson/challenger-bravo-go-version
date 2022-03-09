package task

import (
	"bravo/config"
	"bravo/repository"
	"fmt"
	"time"

	exchange "github.com/asvvvad/exchange"
)

type CurrencyUpdater struct {
	database repository.Database
	config   config.Config
}

var mainServerCurrency string
var amount int
var avaiableCoins []string

func (c *CurrencyUpdater) updateCurrencyTask() {
	ticker := time.NewTicker(5 * time.Second)
	mainServerCurrency = c.config.GetMainServerCurrency()
	currenciesList := c.config.GetInitialCurrencies()
	amount = 1
	for {
		<-ticker.C
		c.updateAllCurrencies(currenciesList)
	}
}

func (c *CurrencyUpdater) updateAllCurrencies(currenciesList []string) {
	for _, currency := range currenciesList {
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

func NewCurrencyManager(database repository.Database, config config.Config) *CurrencyUpdater {
	return &CurrencyUpdater{
		database: database,
		config:   config,
	}
}
