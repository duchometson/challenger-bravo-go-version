package currencyapi

import (
	"math/big"
	"testing"

	"bravo/externalservice/currencyapi/testutil"

	"github.com/asvvvad/exchange"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyApi_ExternalApiGet(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedApi := testutil.NewMockExternalApi(ctrl)

	currencyAPI := New(mockedApi)

	t.Run("testing get from external api with success", func(t *testing.T) {
		expectedResultFromApi := big.NewFloat(1.0)
		first := mockedApi.EXPECT().ValidateCode("BRL").Return(nil)
		mockedApi.EXPECT().LatestRatesSingle("BRL").Return(expectedResultFromApi, nil).After(first)

		value, err := currencyAPI.Get("BRL")
		assert.Nil(t, err)
		assert.Equal(t, float64(1), value)
	})
	t.Run("testing  get from external api when code is invalid", func(t *testing.T) {
		expectedValue := float64(0)
		mockedApi.EXPECT().ValidateCode("BRL").Return(exchange.ErrInvalidCode)

		value, err := currencyAPI.Get("BRL")
		assert.NotNil(t, err)
		assert.Equal(t, exchange.ErrInvalidCode, err)
		assert.Equal(t, expectedValue, value)
	})
	t.Run("testing  get from external api when api returns invalid response", func(t *testing.T) {
		expectedValueFromApi := big.NewFloat(0.0)
		expectedReturn := float64(0)
		first := mockedApi.EXPECT().ValidateCode("BRL").Return(nil)
		mockedApi.EXPECT().LatestRatesSingle("BRL").Return(expectedValueFromApi, exchange.ErrInvalidAPIResponse).After(first)

		value, err := currencyAPI.Get("BRL")
		assert.NotNil(t, err)
		assert.Equal(t, exchange.ErrInvalidAPIResponse, err)
		assert.Equal(t, expectedReturn, value)
	})
}
