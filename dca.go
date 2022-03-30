package dca

import (
	"math"
)

func DeclineCurve(result [][]float64, Qi float64, Di float64, n float64, Dt float64, length int, delay int32) {
	var ai, aim, ai_exp, aim_exp, N float64
	var t_exp_sw, stop int

	if n == 0 {
		ai = -math.Log(1 - (Di / 100))
		aim = ai / 12
	} else if n == 1 {
		N = 0.9998
		ai = ((1 / N) * (math.Pow((1-(Di/100)), -N) - 1))
		aim = ai / 12
		ai_exp = -math.Log(1 - (Dt / 100))
		aim_exp = ai_exp / 12
		t_exp_sw = int((ai/(-math.Log(1-(Dt/100))) - 1) / (ai * N) * 12)
	} else {
		N = n
		ai = ((1 / N) * (math.Pow((1-(Di/100)), -N) - 1))
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

	if n == 0 {
		for i := 0; i < stop; i++ {
			result[i][0] = float64(i)
			result[i][2] = Qi * math.Exp(-aim*result[i][0])
			result[i][3] = Qi * math.Exp(-aim*(result[1][0]+1))
			result[i][1] = (result[i][2] / aim) * (1 - math.Exp(-aim))
			result[i][0] += float64(delay)
		}
	} else {
		for i := 0; i < stop; i++ {
			result[i][0] = float64(i)
			result[i][3] = Qi / math.Pow(1+aim*(result[i][0]+1)*N, 1/N)
			result[i][2] = Qi / math.Pow(1+aim*(result[i][0])*N, 1/N)
			result[i][4] = (1 / N) * math.Pow((1/(1-((result[i][2]-result[i][3])/result[i][2]))), N-1)
			result[i][5] = result[i][4] * 12
			result[i][6] = (1 - (1 / math.Pow(1+result[i][5]*N, 1/N))) * 100
			result[i][1] = ((result[i][2] / (result[i][4] * (1 - N))) * (1 - (1 / math.Pow((1+result[i][4]*N), ((1-N)/N)))))
			result[i][0] += float64(delay)
		}
		for i := stop; i < length; i++ {
			result[i][0] = float64(i)
			result[i][2] = result[int(i)-1][3]
			result[i][3] = result[i][2] * math.Exp(-aim_exp)
			result[i][4] = -math.Log(1 - ((result[i][2] - result[i][3]) / result[i][2]))
			result[i][5] = result[i][4] * 12
			result[i][6] = (1 - math.Exp(-result[i][5])) * 100
			result[i][1] = (result[i][2] / aim_exp) * (1 - math.Exp(-aim_exp))
			result[i][0] += float64(delay)
		}
	}
}
