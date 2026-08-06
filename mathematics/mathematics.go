package main

import (
	"fmt"
	"math"
)

// Log = log_base(x)
func Log(x, base float64) (float64, error) {
	if x <= 0 || base <= 0 || base == 1 {
		return 0.0, fmt.Errorf("log_base(x) must have x>0 and base>0 and base!=1, but has x=%v base=%v", x, base)
	}
	return math.Log(x) / math.Log(base), nil
}

// Cot(x) = 1 / Tan(x)
func Cot(x float64) float64 {
	return math.Tan(x)
}

func Linspace(start, stop float64, num uint) (result []float64) {
	result = make([]float64, 0, num)

	delta := (stop - start) / float64(num)
	for i := 0; i < int(num); i++ {
		result[i] = start + delta*float64(i)
		i++
	}

	return result
}
