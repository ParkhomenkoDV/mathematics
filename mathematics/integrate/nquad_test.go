package integrate

import (
	"math"
	"testing"
)

func TestNQuad(t *testing.T) {
	tests := []struct {
		name string

		f      Func
		ranges []Range

		epsabs float64
		epsrel float64
		limit  int

		args []float64

		want    float64
		tol     float64
		wantErr bool
	}{
		{
			name: "1",
			f: func(x ...float64) float64 {
				return (x[0] - math.Pow(x[0], 1.0/3.0)) / x[0]
			},
			ranges: []Range{
				{1, 8},
			},
			epsabs: 1e-8,
			epsrel: 1e-8,
			limit:  50,
			want:   4.0,
			tol:    1e-6,
		},
		{
			name: "2",
			f: func(x ...float64) float64 {
				return x[0]*x[0] + 2*x[1]
			},
			ranges: []Range{
				{0, 2}, // внешний x
				{0, 1}, // внутренний y
			},
			epsabs: 1e-8,
			epsrel: 1e-8,
			limit:  50,
			want:   14.0 / 3.0,
			tol:    1e-6,
		},
		{
			name: "3",
			f: func(x ...float64) float64 {
				return 2 * x[0] * x[1] * x[1] * x[2]
			},
			ranges: []Range{
				{0, 3},
				{-2, 0},
				{1, 2},
			},
			epsabs: 1e-8,
			epsrel: 1e-8,
			limit:  50,
			want:   36.0,
			tol:    1e-6,
		},
		{
			name: "error: empty ranges",
			f: func(x ...float64) float64 {
				return 0
			},
			ranges:  []Range{},
			wantErr: true,
		},
		{
			name: "error: limit <= 0",
			f: func(x ...float64) float64 {
				return x[0]
			},
			ranges: []Range{
				{0, 1},
			},
			limit:   -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NQuad(tt.f, tt.ranges, tt.args, tt.epsabs, tt.epsrel, tt.limit)
			if tt.wantErr {
				if err == nil {
					t.Fatal("NQuad succeeded unexpectedly")
				}
				return
			}
			if err != nil {
				t.Fatalf("NQuad failed: %v", err)
			}
			if math.Abs(got.Value-tt.want) > tt.tol {
				t.Errorf("NQuad() = %v, want %v (diff %v)", got.Value, tt.want, math.Abs(got.Value-tt.want))
			}
		})
	}
}
