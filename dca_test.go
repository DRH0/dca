package dca

import "testing"

func TestHyperbolic(t *testing.T) {
	type args struct {
		result [][]float64
		Qi     float64
		Di     float64
		n      float64
		Dt     float64
		length int
		delay  int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Hyperbolic(tt.args.result, tt.args.Qi, tt.args.Di, tt.args.n, tt.args.Dt, tt.args.length, tt.args.delay)
		})
	}
}
