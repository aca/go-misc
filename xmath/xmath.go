package xmath

import "math"

func FloorN(f float64, prec int) float64 {
	return math.Floor(f*math.Pow10(prec)) / math.Pow10(prec)
}

func RoundN(f float64, prec int) float64 {
	return math.Round(f*math.Pow10(prec)) / math.Pow10(prec)
}

func CeilN(f float64, prec int) float64 {
	return math.Ceil(f*math.Pow10(prec)) / math.Pow10(prec)
}
