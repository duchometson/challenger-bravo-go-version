package currency

//
//import (
//	"testing"
//
//	"bravo/domain/currency/testutil"
//	"bravo/errorsbravo"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestCurrency(t *testing.T) {
//	ctrl := gomock.NewController(t)
//
//	mockRepository := testutil.NewMockRepository(ctrl)
//
//	currency := New(mockRepository)
//
//	t.Run("testing converting currencies with success", func(t *testing.T) {
//		firstCall := mockRepository.EXPECT().Get("BRL").Return(float64(2), nil)
//
//		mockRepository.EXPECT().Get("USD").Return(float64(10), nil).After(firstCall)
//
//		result, err := currency.Convert("BRL", "USD", 2)
//		assert.Nil(t, err)
//		assert.Equal(t, result, float64(10))
//	})
//	t.Run("testing converting currencies with failed while getting from currency", func(t *testing.T) {
//		mockRepository.EXPECT().Get("BRL").Return(float64(0), errorsbravo.CurrencyDoesntExists)
//
//		result, err := currency.Convert("BRL", "USD", 2)
//		assert.NotNil(t, err)
//		assert.Equal(t, err, errorsbravo.CurrencyDoesntExists)
//		assert.Equal(t, result, float64(0))
//	})
//}
//
