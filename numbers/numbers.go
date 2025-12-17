package numbers

import "math"

func NumberMinusPercent(num int, mdr float64) float64 {
	percent := float64(num) * (mdr / 100)

	delta := float64(num) - percent

	return delta
}

func PercentOf(num int, percent float64) float64 {
	percent = float64(num) * (percent / 100)

	return percent
}

func RevPercentOf(percentage int, percent float64) float64 {
	res := float64(percentage) / (percent / 100)

	return res
}

func RoundToEven(num float64) int {
	return int(math.RoundToEven(num))
}

func Round(num float64) int {
	return int(math.Round(num))
}

// This function is the implementation for BMRI
// The rounding logic
// if the decimal point is greater than or equal to 0.51, then round up (ceil)
// else, round down (floor)
func RoundAtPoint51(data float64) int {
	threshold := float64(0.49)
	return int(math.Floor(data + threshold))
}
