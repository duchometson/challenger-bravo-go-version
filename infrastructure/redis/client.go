package redis

import (
	redis "github.com/go-redis/redis"
)

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

type Client struct {
	redisClient *redis.Client
}

func (c *Client) Get(currency string) (interface{}, error) {
	value, err := c.redisClient.Get(currency).Result()
	if err != nil {
		return 0, c.ErrorNotFound()
	}
	return value, nil
}

func (c *Client) GetAllKeys() ([]string, error) {
	keys, err := c.redisClient.Keys("*").Result()
	if err != nil {
		return []string{}, nil
	}
	return keys, nil
}

func (c *Client) Set(currency string, value interface{}) error {
	err := c.redisClient.Set(currency, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(currency string) error {
	_, err := c.redisClient.Del(currency).Result()
	if err != nil {
		return c.ErrorNotFound()
	}

	return nil
}

func (c *Client) ErrorNotFound() error {
	return redis.TxFailedErr
}

func New() *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &Client{
		redisClient: rdb,
	}
}
