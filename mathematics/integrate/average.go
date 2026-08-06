package integrate

import (
	"errors"
	"math"
)

var ErrZeroDevision = errors.New("zero devision")

// Среднее интегральное.
func Average(f Func, ranges []Range, args []float64, epsabs, epsrel float64, limit int) (float64, error) {
	result, err := NQuad(f, ranges, args, epsabs, epsrel, limit)
	if err != nil {
		return 0.0, err
	}

	var devider = 1.0
	for _, rng := range ranges {
		if rng[0] == rng[1] {
			return math.NaN(), ErrZeroDevision
		}
		devider *= rng[1] - rng[0]
	}

	return result.Value / devider, nil
}
