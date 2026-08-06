package integrate

import (
	"math"
	"testing"
)

func BenchmarkAdaptiveSimpsonPoly(b *testing.B) {
	f := func(x float64) float64 { return x*x + 2*x + 1 }
	rng := Range{0.0, 1.0}
	epsabs, epsrel := 1e-8, 1e-8
	limit := 50
	for i := 0; i < b.N; i++ {
		adaptiveSimpson(f, rng[0], rng[1], epsabs, epsrel, limit)
	}
}

func BenchmarkAdaptiveSimpsonSin(b *testing.B) {
	f := math.Sin
	rng := Range{0.0, math.Pi}
	epsabs, epsrel := 1e-8, 1e-8
	limit := 50
	for i := 0; i < b.N; i++ {
		adaptiveSimpson(f, rng[0], rng[1], epsabs, epsrel, limit)
	}
}
