package integrate

import "testing"

func BenchmarkNQuad1D(b *testing.B) {
	f := func(x ...float64) float64 { return x[0]*x[0] + 1 }
	ranges := []Range{{0, 1}}
	args := []float64{}
	epsabs, epsrel := 1e-8, 1e-8
	limit := 50
	for i := 0; i < b.N; i++ {
		NQuad(f, ranges, args, epsabs, epsrel, limit)
	}
}

func BenchmarkNQuad2D(b *testing.B) {
	f := func(x ...float64) float64 { return x[0] * x[1] }
	ranges := []Range{{0, 1}, {0, 2}}
	args := []float64{}
	epsabs, epsrel := 1e-8, 1e-8
	limit := 50
	for i := 0; i < b.N; i++ {
		NQuad(f, ranges, args, epsabs, epsrel, limit)
	}
}

func BenchmarkNQuad3D(b *testing.B) {
	f := func(x ...float64) float64 { return x[0] * x[1] * x[2] }
	ranges := []Range{{0, 1}, {0, 1}, {0, 1}}
	args := []float64{}
	epsabs, epsrel := 1e-8, 1e-8
	limit := 50
	for i := 0; i < b.N; i++ {
		NQuad(f, ranges, args, epsabs, epsrel, limit)
	}
}
