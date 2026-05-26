package main

import (
	"testing"

	"github.com/user/ton-price-predictor/internal/predictor"
)

func TestPredictorModel(t *testing.T) {
	model := predictor.NewModel(0.95, 0.05)

	tests := []struct {
		current  float64
		expected float64
		trend    string
	}{
		{2.50, 2.425, "down"},
		{10.00, 9.55, "down"},
		{1.00, 1.00, "stable"},
	}

	for _, tt := range tests {
		predicted := model.Predict(tt.current)
		if predicted != tt.expected {
			t.Errorf("Predict(%f) = %f; want %f", tt.current, predicted, tt.expected)
		}
		trend := model.Trend(tt.current, predicted)
		if trend != tt.trend {
			t.Errorf("Trend(%f, %f) = %s; want %s", tt.current, predicted, trend, tt.trend)
		}
	}
}

func TestPredictorStable(t *testing.T) {
	model := predictor.NewModel(1.0, 0.0)
	price := 5.0
	predicted := model.Predict(price)
	if predicted != price {
		t.Errorf("Expected %f, got %f", price, predicted)
	}
	if trend := model.Trend(price, predicted); trend != "stable" {
		t.Errorf("Expected stable, got %s", trend)
	}
}