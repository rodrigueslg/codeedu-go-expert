package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than zero")
}

func TestCalculateTaxWithMock(t *testing.T) {
	repo := &TaxRepositoryMock{}
	repo.On("SaveTax", 10.0).Return(nil).Once()
	repo.On("SaveTax", 0.0).Return(errors.New("amount must be greater than zero"))

	err := CalculateTaxAndSave(1000.0, repo)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0, repo)
	assert.Error(t, err, "amount must be greater than zero")

	repo.AssertExpectations(t)
	repo.AssertNumberOfCalls(t, "SaveTax", 2)
}
