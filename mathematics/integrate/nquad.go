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

	n := len(ranges)
	total := n + len(args)

	// Один срез, который будет использоваться на всех уровнях рекурсии для уменьшения аллокаций
	full := make([]float64, total)
	copy(full[n:], args)

	// Recursive integration function
	var integrate func(depth int) (float64, float64, int, error)
	integrate = func(depth int) (float64, float64, int, error) {
		if depth == len(ranges) {
			return f(full...), 0, 1, nil
		}

		a, b := ranges[depth][0], ranges[depth][1]
		if a == b {
			return 0, 0, 0, nil
		}

		// Define integrand for this dimension: maps x -> integral over remaining dimensions
		integrand := func(x float64) float64 {
			full[depth] = x
			val, _, _, _ := integrate(depth + 1)
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
	val, absErr, nEval, err := integrate(0)
	if err != nil {
		return Result{}, err
	}
	return Result{Value: val, AbsError: absErr, NEval: nEval}, nil
}
