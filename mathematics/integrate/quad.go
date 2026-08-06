package integrate

import (
	"errors"
	"math"
)

// adaptiveSimpson computes the definite integral of f from a to b using adaptive Simpson's rule.
// It returns (integral, error estimate, number of evaluations, error).
func adaptiveSimpson(f func(float64) float64, a, b, epsabs, epsrel float64, limit int) (float64, float64, int, error) {
	if limit <= 0 {
		return 0, 0, 0, errors.New("limit must be positive")
	}
	if a == b {
		return 0, 0, 0, nil
	}
	if a > b {
		a, b = b, a
		// Integrate over swapped interval and negate at end
		val, err, neval, e := adaptiveSimpson(f, a, b, epsabs, epsrel, limit)
		return -val, err, neval, e
	}

	// Recursive helper
	var rec func(a, b float64, fa, fm, fb float64, S, eps float64, depth int) (float64, float64, int, error)
	rec = func(a, b float64, fa, fm, fb float64, S, eps float64, depth int) (float64, float64, int, error) {
		if depth <= 0 {
			return S, math.Abs(S) * eps, 0, errors.New("maximum recursion depth exceeded")
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
			return S2, delta, 3, nil // 3 evaluations: fl, fm, fr (fm already known)
		}
		// Recurse on both halves
		// Evaluate left half
		valL, errL, nevalL, eL := rec(a, mid, fa, fl, fm, leftSimpson, eps/2, depth-1)
		if eL != nil {
			return 0, 0, 0, eL
		}
		// Evaluate right half
		valR, errR, nevalR, eR := rec(mid, b, fm, fr, fb, rightSimpson, eps/2, depth-1)
		if eR != nil {
			return 0, 0, 0, eR
		}
		return valL + valR, errL + errR, nevalL + nevalR + 1, nil // +1 for the fm that was already counted? We'll adjust: neval counts evaluations at new points.
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
	val, err, neval, e := rec(a, b, fa, fm, fb, S, 1.0, depth)
	if e != nil {
		return 0, 0, 0, e
	}
	// Add initial 3 evaluations (fa, fb, fm)
	return val, err, neval + 3, nil
}
