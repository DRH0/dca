package dca

import (
	"math"
)

type ForecastTable struct {
	Qi         float64
	Di         float64
	N          float64
	Dt         float64
	method     string
	length     int
	month      []int
	production []float64
	forecast   []float64
}

//	This gets the initial variables for a procedural decline calculation
func InitialVariables(Qi float64, Di float64, n float64, method string, Dt float64, length int) (float64, float64, float64, int) {
	var ai, aim, ai_exp, aim_exp, N float64
	var t_exp_sw, stop int
	if n == 1 {
		n = 0.9998
	}
	if n == 0 {
		t_exp_sw = 0
		ai = -math.Log(1 - (Di / 100))
		aim = ai / 12
	} else {
		N = n
		if method == "B" {
			ai = ((1 / N) * (math.Pow((1-(Di/100)), -N) - 1))
		} else {
			ai = -math.Log(1 - (Di / 100))
		}
		aim = ai / 12
		ai_exp = -math.Log(1 - (Dt / 100))
		aim_exp = ai_exp / 12
		t_exp_sw = int((ai/(-math.Log(1-(Dt/100))) - 1) / (ai * N) * 12)
	}
	if t_exp_sw > length {
		stop = length
	} else {
		stop = t_exp_sw
	}
	return aim, aim_exp, N, stop
}

// 	Populates a previously allocated 2d array with a decline curve table.
//	If you want exponential decline just use 0 for the n parameter.
//	Table columns:
// 	[index][monthly production][start][end][nom month][nom year][effective %]
func DeclineCurve(result [][]float64, Qi float64, Di float64, n float64, method string, Dt float64, length int) {
	var aim, aim_exp, N, stop = InitialVariables(Qi, Di, n, method, Dt, length)
	//	n=0 means exponential decline from the beginning

	if n == 0 {
		for i := 0; i < stop; i++ {
			result[i][0] = float64(i)
			result[i][2] = Qi * math.Exp(-aim*result[i][0])
			result[i][3] = Qi * math.Exp(-aim*(result[1][0]+1))
			result[i][1] = (result[i][2] / aim) * (1 - math.Exp(-aim))
		}
	} else {
		//	Hyperbolic decline before switch to exponential
		for i := 0; i < stop; i++ {
			result[i][0] = float64(i)
			result[i][3] = Qi / math.Pow(1+aim*(result[i][0]+1)*N, 1/N)
			result[i][2] = Qi / math.Pow(1+aim*(result[i][0])*N, 1/N)
			result[i][4] = (1 / N) * math.Pow((1/(1-((result[i][2]-result[i][3])/result[i][2]))), N-1)
			result[i][5] = result[i][4] * 12
			result[i][6] = (1 - (1 / math.Pow(1+result[i][5]*N, 1/N))) * 100
			result[i][1] = ((result[i][2] / (result[i][4] * (1 - N))) * (1 - (1 / math.Pow((1+result[i][4]*N), ((1-N)/N)))))
		}
		//	exponential decline after switch
		for i := stop; i < length; i++ {
			result[i][0] = float64(i)
			result[i][2] = result[int(i)-1][3]
			result[i][3] = result[i][2] * math.Exp(-aim_exp)
			result[i][4] = -math.Log(1 - ((result[i][2] - result[i][3]) / result[i][2]))
			result[i][5] = result[i][4] * 12
			result[i][6] = (1 - math.Exp(-result[i][5])) * 100
			result[i][1] = (result[i][2] / aim_exp) * (1 - math.Exp(-aim_exp))
		}
	}
}

/*
//  Returns Qi and Di for parameters rolled forward or backwards by n months
func roll(Qi float64, Di float64, n float64, Dt float64, offset int) []float64 {
	var length, i, j int
	if offset < 0 {
		length = -offset
	} else {
		length = offset + 1
	}
	table := prz.Make2df64(10, 7)
	result := make([]float64, 2)
	// Put some error checking here, check Dt>Di, etc.
	var aim, aim_exp, N, stop = InitialVariables(Qi, Di, n, Dt, length)

}*/

//	Converts B to H
func BtoH(n_b float64, d_b float64) float64 {
	var ai float64 = (1 / n_b * (math.Pow(1-d_b/100, -n_b) - 1) / 12)
	return (1 - math.Exp(-12*(ai))) * 100
}

//	Converts H to B
func HtoB(n_h float64, d_h float64) float64 {
	var ai float64 = -math.Log(1-d_h/100) / 12
	return (1 - math.Pow(1/(12*ai*n_h+1), 1/n_h)) * 100
}
