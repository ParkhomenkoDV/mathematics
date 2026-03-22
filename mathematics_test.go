package main

import (
	"math"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{
			name:    "a<0",
			a:       -1,
			b:       10,
			want:    math.Log(10) / math.Log(-1),
			wantErr: true,
		}, {
			name:    "a=1",
			a:       1,
			b:       10,
			want:    math.Log(10) / math.Log(1),
			wantErr: true,
		}, {
			name:    "b<0",
			a:       2,
			b:       -1,
			want:    math.Log(-1) / math.Log(2),
			wantErr: true,
		}, {
			name:    "ok",
			a:       2,
			b:       10,
			want:    math.Log(10) / math.Log(2),
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, gotErr := Log(test.a, test.b)
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
