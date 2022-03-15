package currency

import (
	"time"
)

type Worker struct {
	currencyService Service
	externalService ExternalService
}

var DefaultSupportedCurrencies = []string{"BTC", "BRL", "ETH", "EUR", "USD"}

func (w *Worker) Update() {
	w.UpdateCurrencies(DefaultSupportedCurrencies)

	for {
		currencies, err := w.currencyService.GetAllKeys()
		if err != nil {
			break
		}

		w.UpdateCurrencies(currencies)

		time.Sleep(5 * time.Second)
	}
}

func (w *Worker) UpdateCurrencies(currencies []string) {
	for _, currency := range currencies {
		value, _ := w.externalService.Get(currency)
		w.currencyService.Set(currency, value)
	}
}

func NewWorker(currencyService Service, externalService ExternalService) *Worker {
	return &Worker{
		currencyService: currencyService,
		externalService: externalService,
	}
}
