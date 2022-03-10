package redis

import (
	"bravo/errorsbravo"

	"github.com/go-redis/redis"
)

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

type Client struct {
	redisClient *redis.Client
}

// TO IMPLEMENT USING REDIS:
//Get(string) (interface{}, error)
//Set(string, interface{}) error
//Delete(string) error
//ErrorNotFound() error

// https://github.com/go-redis/redis

func (m *Client) Get(currency string) (interface{}, error) {
	redisCmd := m.redisClient.Get(currency)
	if redisCmd != nil {
		return 0, errorsbravo.CURRENCY_DOESNT_EXISTS
	}

	return redisCmd.Val(), nil
}

func (m *Client) GetAllKeys() ([]string, error) {
	return []string{}, nil
}

func (m *Client) Set(currency string, value interface{}) error {
	m.redisClient.Set("currency", value, 0)
	return nil
}

func (m *Client) Delete(currency string) error {
	redisCmd := m.redisClient.Del(currency)
	if redisCmd != nil {
		return errorsbravo.CURRENCY_DOESNT_EXISTS
	}

	return nil
}

func (m *Client) ErrorNotFound() error {
	return nil
}

func New() *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Client{
		redisClient: rdb,
	}
}
