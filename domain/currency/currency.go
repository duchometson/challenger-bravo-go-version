package currency

type Currency struct {
	repository Repository
}

func (c *Currency) Get(currency string) (float64, error) {
	return c.repository.Get(currency)
}

func (c *Currency) Convert(from, to string, amount float64) (float64, error) {

	return 0, nil
}

// TO IMPLEMENT THE REMAINING METHODS

func New(repository Repository) *Currency {
	return &Currency{
		repository: repository,
	}
}
