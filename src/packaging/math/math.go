package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Max(xs []float64) float64 {
	var ma float64
	for i, x := range xs {
		if i == 0 || x > ma {
			ma = x
		}
	}
	return ma
}

func Min(xs []float64) float64 {
	var mi float64
	for i, x := range xs {
		if i == 0 || x < mi {
			mi = x
		}
	}
	return mi
}
