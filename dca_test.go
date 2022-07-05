package dca

import (
	"fmt"
	"testing"
)

//	these functions are in prz
func Make2df64(rows int, cols int) [][]float64 {
	table := make([][]float64, rows)
	chunk := make([]float64, cols*rows)
	for i := range table {
		table[i], chunk = chunk[:cols], chunk[cols:]
	}
	return table
}
func Print2df64(table [][]float64) {
	for _, row := range table {
		for _, col := range row {
			fmt.Print(fmt.Sprintf("%.2f", col) + " ")
		}
		fmt.Print("\n")
	}
}
func TestDeclineCurve(t *testing.T) {

	var tests = []struct {
		name   string
		Qi     float64
		Di     float64
		n      float64
		Dt     float64
		length int
		delay  int32
		result [][]float64
	}{
		{"Test 1", 25000, 80, 1.2, 7, 10, 0, Make2df64(10, 7)},
		{"Test 2", 25000, 0.01, 0, 7, 10, 0, Make2df64(10, 7)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeclineCurve(tt.result, tt.Qi, tt.Di, tt.n, tt.Dt, tt.length, tt.delay)
			//Print2df64(tt.result)
		})
	}
}
