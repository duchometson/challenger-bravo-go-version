package currency

type Currency struct {
	repository Repository
}

func (c *Currency) Get(currency string) (float64, error) {
	return c.repository.Get(currency)
}

func (c *Currency) GetAllKeys() ([]string, error) {
	return c.repository.GetAllKeys()
}

func (c *Currency) Convert(from, to string, amount float64) (float64, error) {
	valueFrom, err := c.repository.Get(from)
	if err != nil {
		return 0, err
	}

	valueTo, err := c.repository.Get(to)
	if err != nil {
		return 0, err
	}

	convertedValue := (valueTo / valueFrom) * amount

	return convertedValue, nil
}

func (c *Currency) Set(currency string, value float64) error {
	return c.repository.Set(currency, value)
}

func (c *Currency) Delete(currency string) error {
	return c.repository.Delete(currency)
}

func New(repository Repository) *Currency {
	return &Currency{
		repository: repository,
	}
}
