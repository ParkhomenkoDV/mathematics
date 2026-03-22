package main

import (
	"fmt"
	"math"
)

// Log = log_a(b)
func Log(a, b float64) (float64, error) {
	if a <= 0 || a == 1 || b <= 0 {
		return 0.0, fmt.Errorf("log_a(b) must have a>0 and a!=1 and b>0, but has a=%v b=%v", a, b)
	}
	return math.Log(b) / math.Log(a), nil
}

// Cot(x) = 1 / Tan(x)
func Cot(x float64) float64 {
	return math.Tan(x)
}
