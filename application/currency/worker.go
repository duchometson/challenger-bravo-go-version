package currency

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Worker struct {
	currencyService            Service
	defaultSupportedCurrencies []string
	externalService            ExternalService
	interval                   time.Duration
}

func (w *Worker) Update() {
	w.UpdateCurrencies(w.defaultSupportedCurrencies)

	for {
		currencies, err := w.currencyService.GetAllKeys()
		if err != nil {
			break
		}

		w.UpdateCurrencies(currencies)

		time.Sleep(w.interval)
	}
}

func (w *Worker) UpdateCurrencies(currencies []string) {
	for _, currency := range currencies {
		value, err := w.externalService.Get(currency)
		if err != nil {
			log.WithError(err).WithField("currency", currency).Error("failed to get currency")
			continue
		}

		if err := w.currencyService.Set(currency, value); err != nil {
			log.WithError(err).WithField("currency", currency).Error("failed to set currency")
			continue
		}

		log.WithField("currency", currency).Info("currency updated successfully")
	}
}

func NewWorker(currencyService Service, externalService ExternalService, interval time.Duration, defaultSupportedCurrencies []string) *Worker {
	return &Worker{
		currencyService:            currencyService,
		defaultSupportedCurrencies: defaultSupportedCurrencies,
		externalService:            externalService,
		interval:                   interval,
	}
}
