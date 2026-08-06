package main

import (
	"math"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		name    string
		base    float64
		x       float64
		want    float64
		wantErr bool
	}{
		{
			name:    "a<0",
			base:    -1,
			x:       10,
			want:    math.Log(10) / math.Log(-1),
			wantErr: true,
		}, {
			name:    "a=1",
			base:    1,
			x:       10,
			want:    math.Log(10) / math.Log(1),
			wantErr: true,
		}, {
			name:    "b<0",
			base:    2,
			x:       -1,
			want:    math.Log(-1) / math.Log(2),
			wantErr: true,
		}, {
			name:    "ok",
			base:    2,
			x:       10,
			want:    math.Log(10) / math.Log(2),
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotErr := Log(test.x, test.base)
			if gotErr != nil {
				if !test.wantErr {
					t.Errorf("Log() failed: %v", gotErr)
				}
				return
			}
			if test.wantErr {
				t.Fatal("Log() succeeded unexpectedly")
			}
			if got != test.want {
				t.Errorf("Log() = %v, want %v", got, test.want)
			}
		})
	}
}
