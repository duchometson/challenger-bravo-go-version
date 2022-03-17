package currency

import (
	"testing"

	"bravo/domain/currency/testutil"
	"bravo/errorsbravo"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCurrency_CurrencyConversion(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := testutil.NewMockRepository(ctrl)

	currency := New(mockRepository)

	t.Run("testing converting currencies with success", func(t *testing.T) {
		firstCall := mockRepository.EXPECT().Get("BRL").Return(float64(2), nil)

		mockRepository.EXPECT().Get("USD").Return(float64(10), nil).After(firstCall)

		result, err := currency.Convert("BRL", "USD", 2)
		assert.Nil(t, err)
		assert.Equal(t, result, float64(10))
	})
	t.Run("testing converting currencies when 'currency to' fails while getting from currency", func(t *testing.T) {
		mockRepository.EXPECT().Get("BRL").Return(float64(0), errorsbravo.CurrencyDoesntExists)

		result, err := currency.Convert("BRL", "USD", 2)
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.CurrencyDoesntExists)
		assert.Equal(t, result, float64(0))
	})
	t.Run("testing converting currencies when 'currency from' fails while getting from currency", func(t *testing.T) {
		firstCall := mockRepository.EXPECT().Get("BRL").Return(float64(2), nil)

		mockRepository.EXPECT().Get("USD").Return(float64(0), errorsbravo.CurrencyDoesntExists).After(firstCall)

		result, err := currency.Convert("BRL", "USD", 2)
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.CurrencyDoesntExists)
		assert.Equal(t, result, float64(0))
	})
}

func TestCurrency_CurrencyGet(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := testutil.NewMockRepository(ctrl)

	currency := New(mockRepository)

	t.Run("testing getting currencies with success", func(t *testing.T) {
		mockRepository.EXPECT().Get("BRL").Return(float64(2), nil)
		result, err := currency.Get("BRL")
		assert.Nil(t, err)
		assert.Equal(t, result, float64(2))
	})
	t.Run("testing getting currencies that doesnt exists", func(t *testing.T) {
		mockRepository.EXPECT().Get("BRL").Return(float64(0), errorsbravo.CurrencyDoesntExists)
		result, err := currency.Get("BRL")
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.CurrencyDoesntExists)
		assert.Equal(t, result, float64(0))
	})
}

func TestCurrency_CurrencySet(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := testutil.NewMockRepository(ctrl)

	currency := New(mockRepository)

	t.Run("testing setting currencies with success", func(t *testing.T) {
		mockRepository.EXPECT().Set("BRL", 1.0).Return(nil)
		err := currency.Set("BRL", 1.0)
		assert.Nil(t, err)
	})
	t.Run("testing setting currencies but receives internal error", func(t *testing.T) {
		mockRepository.EXPECT().Set("BRL", 1.0).Return(errorsbravo.InternalError)
		err := currency.Set("BRL", 1.0)
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.InternalError)
	})
}

func TestCurrency_CurrencyDelete(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := testutil.NewMockRepository(ctrl)

	currency := New(mockRepository)

	t.Run("testing deletting currencies with success", func(t *testing.T) {
		mockRepository.EXPECT().Delete("BRL").Return(nil)
		err := currency.Delete("BRL")
		assert.Nil(t, err)
	})
	t.Run("testing deletting currencies but currencies doesnt exists", func(t *testing.T) {
		mockRepository.EXPECT().Delete("aaaa").Return(errorsbravo.CurrencyDoesntExists)
		err := currency.Delete("aaaa")
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.CurrencyDoesntExists)
	})
}

func TestCurrency_CurrencyGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := testutil.NewMockRepository(ctrl)

	currency := New(mockRepository)

	t.Run("testing getting all currencies with success", func(t *testing.T) {
		expectedKeysList := []string{"BRL", "USD"}
		mockRepository.EXPECT().GetAllKeys().Return(expectedKeysList, nil)
		resultList, err := currency.GetAllKeys()
		assert.Nil(t, err)
		assert.Equal(t, resultList, expectedKeysList)
	})
	t.Run("testing getting all currencies but internal error occurrs", func(t *testing.T) {
		mockRepository.EXPECT().GetAllKeys().Return([]string{}, errorsbravo.InternalError)
		resultList, err := currency.GetAllKeys()
		assert.NotNil(t, err)
		assert.Equal(t, err, errorsbravo.InternalError)
		assert.Equal(t, resultList, []string{})
	})
}
