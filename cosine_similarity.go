package cosinesimilarity

import "math"

func Compute(a, b [][]float64) [][]float64 {
	result := make([][]float64, len(a))
	for iX, x := range a {
		result[iX] = make([]float64, len(b))
		for iY, y := range b {
			result[iX][iY] = compute(x, y)
		}
	}
	return result
}

func compute(x, y []float64) float64 {
	var sum, s1, s2 float64
	for i := 0; i < len(x); i++ {
		sum += x[i] * y[i]
		s1 += math.Pow(x[i], 2)
		s2 += math.Pow(y[i], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0
	}
	return sum / (math.Sqrt(s1) * math.Sqrt(s2))
}
