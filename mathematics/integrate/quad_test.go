package integrate

import (
	"math"
	"testing"
)

func TestAdaptiveSimpson(t *testing.T) {
	tests := []struct {
		name   string
		f      func(float64) float64
		a, b   float64
		epsabs float64
		epsrel float64
		limit  int
		want   float64
		tol    float64
	}{
		{
			name:   "x^2 on [0,1]",
			f:      func(x float64) float64 { return x * x },
			a:      0,
			b:      1,
			epsabs: 1e-8,
			epsrel: 1e-8,
			limit:  50,
			want:   1.0 / 3.0,
			tol:    1e-6,
		},
		{
			name:   "sin(x) on [0,pi]",
			f:      math.Sin,
			a:      0,
			b:      math.Pi,
			epsabs: 1e-8,
			epsrel: 1e-8,
			limit:  50,
			want:   2.0,
			tol:    1e-6,
		},
		{
			name:   "zero interval",
			f:      func(x float64) float64 { return x },
			a:      2,
			b:      2,
			epsabs: 1e-8,
			epsrel: 1e-8,
			limit:  50,
			want:   0,
			tol:    0,
		},
		{
			name:   "negative interval (a>b)",
			f:      func(x float64) float64 { return x },
			a:      1,
			b:      0,
			epsabs: 1e-8,
			epsrel: 1e-8,
			limit:  50,
			want:   -0.5,
			tol:    1e-6,
		},
		{
			name:   "constant 5 on [0,2]",
			f:      func(x float64) float64 { return 5 },
			a:      0,
			b:      2,
			epsabs: 1e-8,
			epsrel: 1e-8,
			limit:  50,
			want:   10,
			tol:    1e-6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, _, err := adaptiveSimpson(tt.f, tt.a, tt.b, tt.epsabs, tt.epsrel, tt.limit)
			if err != nil {
				t.Fatalf("adaptiveSimpson failed: %v", err)
			}
			if math.Abs(got-tt.want) > tt.tol {
				t.Errorf("adaptiveSimpson() = %v, want %v (diff %v)", got, tt.want, math.Abs(got-tt.want))
			}
		})
	}
}
