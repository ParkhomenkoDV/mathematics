package integrate

import (
	"errors"
)

var ErrRandes = errors.New("at least one range must be provided")

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
	epsAbs, epsRel float64,
	limit int,
) (Result, error) {
	// Validate ranges
	if len(ranges) == 0 {
		return Result{}, ErrRandes
	}

	n := len(ranges)
	total := n + len(args)

	// Один срез, который будет использоваться на всех уровнях рекурсии для уменьшения аллокаций
	full := make([]float64, total)
	copy(full[n:], args)

	// Recursive integration function
	var integrate func(depth int) (Result, error)
	integrate = func(depth int) (Result, error) {
		if depth == len(ranges) {
			return Result{Value: f(full...), AbsError: 0, NEval: 1}, nil
		}

		a, b := ranges[depth][0], ranges[depth][1]
		if a == b {
			return Result{Value: 0, AbsError: 0, NEval: 0}, nil
		}

		// Define integrand for this dimension: maps x -> integral over remaining dimensions
		integrand := func(x float64) float64 {
			full[depth] = x
			result, _ := integrate(depth + 1)
			return result.Value
		}

		// Perform 1D adaptive integration
		result, err := adaptiveSimpson(integrand, a, b, epsAbs, epsRel, limit)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	// Start recursion
	result, err := integrate(0)
	if err != nil {
		return result, err
	}
	return result, nil
}
