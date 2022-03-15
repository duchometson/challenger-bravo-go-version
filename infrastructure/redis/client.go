package redis

import (
	"context"
	redis "github.com/go-redis/redis/v8"
)

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

type Client struct {
	redisClient *redis.Client
}

func (c *Client) Get(ctx context.Context, currency string) (interface{}, error) {
	value, err := c.redisClient.Get(ctx, currency).Result()
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
	if err := c.redisClient.Set(currency, value, 0).Err(); err != nil {
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

func New(address, password string, db int) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	return &Client{
		redisClient: rdb,
	}
}
