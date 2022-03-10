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

	return 0, nil
}

func (c *Currency) Set(currency string, value float64) error {
	return c.repository.Set(currency, value)
}

func (c *Currency) Delete(currency string) error {
	return c.repository.Delete(currency)
}

// TO IMPLEMENT THE REMAINING METHODS

func New(repository Repository) *Currency {
	return &Currency{
		repository: repository,
	}
}
