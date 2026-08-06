package integrate

import (
	"errors"
)

// NQuad computes a multiple integral over several variables.
//
// Parameters:
//   - f: integrand function of the form f(x0, x1, ..., xn, args...).
//   - ranges: slice of range definitions. Each element is either [2]float64 for fixed bounds,
//     or a Range that returns bounds based on outer variables and args.
//   - args: extra arguments passed to f and range functions.
//
// Returns:
//   - Result containing integral estimate, error estimate, and number of evaluations, and an error if any.
func NQuad(
	f Func,
	ranges []Range,
	args []float64,
	epsabs, epsrel float64,
	limit int,
) (Result, error) {
	// Validate ranges
	if len(ranges) == 0 {
		return Result{}, errors.New("at least one range must be provided")
	}

	// Recursive integration function
	var integrate func(depth int, params []float64) (float64, float64, int, error)
	integrate = func(depth int, params []float64) (float64, float64, int, error) {
		if depth == len(ranges) {
			// Base case: evaluate integrand at current point
			fullArgs := append(params, args...)
			val := f(fullArgs...)
			return val, 0, 1, nil
		}

		// Determine bounds for this dimension
		var a, b float64
		rng := ranges[depth]
		a, b = rng[0], rng[1]

		// If limits are equal, integral is zero
		if a == b {
			return 0, 0, 0, nil
		}

		// Define integrand for this dimension: maps x -> integral over remaining dimensions
		integrand := func(x float64) float64 {
			newParams := append(params, x)
			val, _, _, _ := integrate(depth+1, newParams)
			return val
		}

		// Perform 1D adaptive integration
		res, err, neval, err2 := adaptiveSimpson(integrand, a, b, epsabs, epsrel, limit)
		if err2 != nil {
			return 0, 0, 0, err2
		}
		// Note: error returned is absolute error estimate; we keep it as is.
		return res, err, neval, nil
	}

	// Start recursion
	val, err, neval, e := integrate(0, []float64{})
	if e != nil {
		return Result{}, e
	}
	return Result{Value: val, AbsError: err, NEval: neval}, nil
}
