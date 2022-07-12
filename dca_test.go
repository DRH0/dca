package dca

import (
	"testing"

	"github.com/DRH0/prz"
)

func TestDeclineCurve(t *testing.T) {

	var tests = []struct {
		name   string
		Qi     float64
		Di     float64
		n      float64
		Dt     float64
		length int
		result [][]float64
	}{
		{"Test 1", 25000, 80, 1.2, 7, 10, prz.Make2df64(10, 7)},
		{"Test 2", 25000, 80, 1.2, 7, 10, prz.Make2df64(10, 7)},
		{"Test 3", 25000, 0.01, 0, 7, 10, prz.Make2df64(10, 7)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeclineCurve(tt.result, tt.Qi, tt.Di, tt.n, tt.Dt, tt.length)
			prz.Print2df64(tt.result)
		})
	}
}
