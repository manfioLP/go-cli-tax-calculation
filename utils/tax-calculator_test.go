package utils

import (
	"testing"
)

func TestCalculateWeightedPrice(t *testing.T) {
	w := calculateWeightedAverage(0.0, 0, 20.00, 10)
	expected := 20.0
	if eq := w == expected; !eq {
		t.Errorf("Expected weighted average to be %f but got %f", expected, w)
	}

	w = calculateWeightedAverage(20.0, 5, 10.00, 5)
	expected = 15.0
	if eq := w == expected; !eq {
		t.Errorf("Expected weighted average to be %f but got %f", expected, w)
	}
}

func TestCalculateWeightedPrice2(t *testing.T) {
	w := calculateWeightedAverage(0.0, 0, 10.00, 100)
	expected := 10.0
	if eq := w == expected; !eq {
		t.Errorf("Expected weighted average to be %f but got %f", expected, w)
	}

	w = calculateWeightedAverage(w, 0, 10.00, 10000)
	expected = 10.0
	if eq := w == expected; !eq {
		t.Errorf("Expected weighted average to be %f but got %f", expected, w)
	}
}

func TestCalculateTax(t *testing.T) {
	// Test case 1
	input1 := []Order{
		{Operation: "buy", UnitCost: 10.00, Quantity: 100},
		{Operation: "sell", UnitCost: 15.00, Quantity: 50},
		{Operation: "sell", UnitCost: 15.00, Quantity: 50},
	}
	expected1 := []TaxResponse{
		{Tax: 0.00},
		{Tax: 0.00},
		{Tax: 0.00},
	}
	result1 := CalculateTax(input1)
	compareResults(t, expected1, result1)

	// Test case 2
	input2 := []Order{
		{Operation: "buy", UnitCost: 10.00, Quantity: 100},
		{Operation: "sell", UnitCost: 15.00, Quantity: 50},
		{Operation: "sell", UnitCost: 15.00, Quantity: 50},
		{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
		{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
	}
	expected2 := []TaxResponse{
		{Tax: 0.00},
		{Tax: 0.00},
		{Tax: 0.00},
		{Tax: 0.00},
		{Tax: 10000.00},
	}
	result2 := CalculateTax(input2)
	compareResults(t, expected2, result2)
}

func compareResults(t *testing.T, expected, result []TaxResponse) {
	if len(expected) != len(result) {
		t.Errorf("Expected %d operations, but got %d", len(expected), len(result))
		return
	}

	for i := 0; i < len(expected); i++ {
		eq := expected[i].Tax == result[i].Tax
		if !eq {
			t.Errorf("At index %d expected %f but got %f", i, expected[i].Tax, result[i].Tax)
		}
	}
}
