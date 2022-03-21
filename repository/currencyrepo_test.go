package repository

import (
	"bravo/errorsbravo"
	"bravo/repository/testutil"
	"context"
	"testing"

	"github.com/go-redis/redis"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyRepository_CurrencyGet(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockDatabase := testutil.NewMockDatabase(ctrl)

	currency := New(mockDatabase)

	t.Run("testing get from repository with success", func(t *testing.T) {
		expectedReturn := float64(1)
		expectedDatabaseReturn := string("1.0")
		ctx := context.TODO()

		mockDatabase.EXPECT().Get(ctx, "BRL").Return(expectedDatabaseReturn, nil)

		value, err := currency.Get("BRL")
		assert.Nil(t, err)
		assert.Equal(t, expectedReturn, value)
	})

	t.Run("testing get from repository when currency doesnt exists error is returned", func(t *testing.T) {
		expectedReturn := float64(0.0)
		ctx := context.TODO()
		first := mockDatabase.EXPECT().Get(ctx, "BRL").Return(nil, redis.TxFailedErr)
		mockDatabase.EXPECT().ErrorNotFound().Return(redis.TxFailedErr).After(first)

		value, err := currency.Get("BRL")
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.CurrencyDoesntExists)
		assert.Equal(t, expectedReturn, value)
	})
	t.Run("testing get from repository when internal error returned from database", func(t *testing.T) {
		expectedReturn := float64(0.0)
		ctx := context.TODO()
		first := mockDatabase.EXPECT().Get(ctx, "BRL").Return(nil, errorsbravo.InternalError)
		mockDatabase.EXPECT().ErrorNotFound().Return(redis.TxFailedErr).After(first)

		value, err := currency.Get("BRL")
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.InternalError)
		assert.Equal(t, expectedReturn, value)
	})
	t.Run("testing get from repository when internal error is returned from parsing", func(t *testing.T) {
		expectedDatabaseReturn := string("abc")
		expectedReturn := float64(0.0)
		ctx := context.TODO()
		mockDatabase.EXPECT().Get(ctx, "BRL").Return(expectedDatabaseReturn, nil)

		value, err := currency.Get("BRL")
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.InternalError)
		assert.Equal(t, expectedReturn, value)
	})
}

func TestCurrencyRepository_CurrencyGetAllKeys(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockDatabase := testutil.NewMockDatabase(ctrl)

	currency := New(mockDatabase)

	t.Run("testing getall from repository with success", func(t *testing.T) {
		expectedDatabaseReturn := []string{"BRL", "EUR"}
		ctx := context.TODO()

		mockDatabase.EXPECT().GetAllKeys(ctx).Return(expectedDatabaseReturn, nil)

		value, err := currency.GetAllKeys()
		assert.Nil(t, err)
		assert.Equal(t, expectedDatabaseReturn, value)
	})
	t.Run("testing getall from repository when currency doesnt exists error is returned", func(t *testing.T) {
		expectedDatabaseReturn := []string{}
		ctx := context.TODO()

		first := mockDatabase.EXPECT().GetAllKeys(ctx).Return(expectedDatabaseReturn, redis.TxFailedErr)
		mockDatabase.EXPECT().ErrorNotFound().Return(redis.TxFailedErr).After(first)

		value, err := currency.GetAllKeys()
		assert.NotNil(t, err, errorsbravo.CurrencyDoesntExists)
		assert.Equal(t, expectedDatabaseReturn, value)
	})
	t.Run("testing getall from repository when internal error is returned", func(t *testing.T) {
		expectedDatabaseReturn := []string{}
		ctx := context.TODO()

		first := mockDatabase.EXPECT().GetAllKeys(ctx).Return(expectedDatabaseReturn, errorsbravo.InternalError)
		mockDatabase.EXPECT().ErrorNotFound().Return(redis.TxFailedErr).After(first)

		value, err := currency.GetAllKeys()
		assert.NotNil(t, err, errorsbravo.CurrencyDoesntExists)
		assert.Equal(t, expectedDatabaseReturn, value)
	})
}

func TestCurrencyRepository_CurrencySet(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockDatabase := testutil.NewMockDatabase(ctrl)

	currency := New(mockDatabase)

	t.Run("testing set from repository with success", func(t *testing.T) {
		ctx := context.TODO()

		mockDatabase.EXPECT().Set(ctx, "BRL", float64(1.0)).Return(nil)

		err := currency.Set("BRL", float64(1.0))
		assert.Nil(t, err)
	})
	t.Run("testing set from repository when internal error is returned", func(t *testing.T) {
		ctx := context.TODO()

		mockDatabase.EXPECT().Set(ctx, "BRL", float64(1.0)).Return(redis.TxFailedErr)

		err := currency.Set("BRL", float64(1.0))
		assert.NotNil(t, err, errorsbravo.InternalError)
	})
}

func TestCurrencyRepository_CurrencyDelete(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockDatabase := testutil.NewMockDatabase(ctrl)

	currency := New(mockDatabase)

	t.Run("testing delete from repository with success", func(t *testing.T) {
		ctx := context.TODO()

		mockDatabase.EXPECT().Delete(ctx, "BRL").Return(nil)

		err := currency.Delete("BRL")
		assert.Nil(t, err)
	})
	t.Run("testing delete from repository when currency doesnt exists error is returned", func(t *testing.T) {
		ctx := context.TODO()

		first := mockDatabase.EXPECT().Delete(ctx, "BRL").Return(redis.TxFailedErr)
		mockDatabase.EXPECT().ErrorNotFound().Return(redis.TxFailedErr).After(first)

		err := currency.Delete("BRL")
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.CurrencyDoesntExists)
	})
	t.Run("testing delete from repository when internal error is returned", func(t *testing.T) {
		ctx := context.TODO()

		first := mockDatabase.EXPECT().Delete(ctx, "BRL").Return(errorsbravo.InternalError)
		mockDatabase.EXPECT().ErrorNotFound().Return(redis.TxFailedErr).After(first)

		err := currency.Delete("BRL")
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.InternalError)
	})
}
