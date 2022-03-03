package service

import "bravo/model"

type Converter struct {
	database Database
}

func (c *Converter) Convert(from string, to string, value float64) (float64, error) {
	fromValue, err := c.database.Get(from)
	if err != nil {
		return 0, model.NewApplicationError(err, from)
	}

	toValue, err := c.database.Get(to)
	if err != nil {
		return 0, model.NewApplicationError(err, to)
	}

	coefficientOfConversion := fromValue / toValue

	return coefficientOfConversion * value, nil
}

func NewConverter(database Database) *Converter {
	return &Converter{
		database: database,
	}
}
