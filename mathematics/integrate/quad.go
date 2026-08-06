package integrate

import (
	"errors"
	"math"
)

var (
	ErrNegativeLimit = errors.New("limit must be positive")
	ErrMaxRecursion  = errors.New("maximum recursion depth exceeded")
)

// adaptiveSimpson computes the definite integral of f from a to b using adaptive Simpson's rule.
// It returns (integral, error estimate, number of evaluations, error).
func adaptiveSimpson(f func(float64) float64, a, b, epsabs, epsrel float64, limit int) (Result, error) {
	if limit <= 0 {
		return Result{}, ErrNegativeLimit
	}
	if a == b {
		return Result{Value: 0, AbsError: 0, NEval: 0}, nil
	}
	if a > b {
		a, b = b, a
		// Integrate over swapped interval and negate at end
		result, err := adaptiveSimpson(f, a, b, epsabs, epsrel, limit)
		result.Value = -result.Value
		return result, err
	}

	// Recursive helper
	var rec func(a, b float64, fa, fm, fb float64, S, eps float64, depth int) (Result, error)
	rec = func(a, b float64, fa, fm, fb float64, S, eps float64, depth int) (Result, error) {
		if depth <= 0 {
			return Result{Value: S, AbsError: math.Abs(S) * eps, NEval: 0}, ErrMaxRecursion
		}
		// Simpson's rule on whole interval
		// S = (b-a)/6 * (fa + 4*fm + fb)
		// Midpoint of left and right
		mid := (a + b) / 2
		// Points for subintervals
		leftMid, rightMid := (a+mid)/2, (mid+b)/2
		fl, fr := f(leftMid), f(rightMid)
		// Simpson on left and right
		leftSimpson := (mid - a) / 6 * (fa + 4*fl + fm)
		rightSimpson := (b - mid) / 6 * (fm + 4*fr + fb)
		// Combined Simpson
		S2 := leftSimpson + rightSimpson
		// Error estimate: |S2 - S|/15
		delta := math.Abs(S2-S) / 15.0
		// Tolerance
		tol := epsabs + epsrel*math.Abs(S2)
		if delta <= tol || depth == 1 {
			// convergence or forced stop
			return Result{Value: S2, AbsError: delta, NEval: 3}, nil // 3 evaluations: fl, fm, fr (fm already known)
		}
		// Recurse on both halves
		// Evaluate left half
		resultL, errL := rec(a, mid, fa, fl, fm, leftSimpson, eps/2, depth-1)
		if errL != nil {
			return Result{}, errL
		}
		// Evaluate right half
		resultR, errR := rec(mid, b, fm, fr, fb, rightSimpson, eps/2, depth-1)
		if errR != nil {
			return Result{}, errR
		}
		return Result{
			Value:    resultL.Value + resultR.Value,
			AbsError: resultL.AbsError + resultR.AbsError,
			NEval:    resultL.NEval + resultR.NEval + 1, // +1 for the fm that was already counted? We'll adjust: neval counts evaluations at new points.
		}, nil
	}

	// Initial evaluations
	fa, fb := f(a), f(b)
	fm := f((a + b) / 2)
	S := (b - a) / 6 * (fa + 4*fm + fb)
	// Set depth based on limit: we use limit as max number of subdivisions? Actually limit in scipy is max subintervals.
	// We'll use limit as max depth (number of bisections). We'll map limit to depth roughly log2(limit)+1.
	depth := int(math.Ceil(math.Log2(float64(limit)))) + 2
	if depth < 1 {
		depth = 1
	}
	result, err := rec(a, b, fa, fm, fb, S, 1.0, depth)
	if err != nil {
		return result, err
	}
	// Add initial 3 evaluations (fa, fb, fm)
	result.NEval += 3
	return result, nil
}
