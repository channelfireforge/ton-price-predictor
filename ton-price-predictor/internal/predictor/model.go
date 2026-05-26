package predictor

import "math"

type Model struct {
	Alpha float64
	Beta  float64
}

func NewModel(alpha, beta float64) *Model {
	return &Model{Alpha: alpha, Beta: beta}
}

func (m *Model) Predict(currentPrice float64) float64 {
	return currentPrice*m.Alpha + m.Beta
}

func (m *Model) Trend(current, predicted float64) string {
	diff := predicted - current
	if math.Abs(diff) < 0.01 {
		return "stable"
	} else if diff > 0 {
		return "up"
	}
	return "down"
}