package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MockShopModel struct{}

func (m *MockShopModel) CountCustomers(_ time.Time) (int, error) {
	return 1000, nil
}

func (m *MockShopModel) CountSales(_ time.Time) (int, error) {
	return 333, nil
}

func TestCalculateSalesRate(t *testing.T) {
	// Initialize the mock.
	msm := &MockShopModel{}
	// Pass the mock to the calculateSalesRate() function.
	sr, err := calculateSalesRate(msm)
	assert.NoError(t, err)

	// Check that the return value is as expected, based on the mocked inputs
	assert.Equal(t, "0.33", sr)
}
