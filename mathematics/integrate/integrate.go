// Package integrate provides numerical integration routines.
package integrate

const (
	EpsAbs = 1.49e-8
	EpsRel = 1.49e-8
	Limit  = 50
)

type Func func(...float64) float64

// Range is a function that returns the lower and upper bounds for a given dimension
// based on the current values of outer integration variables and additional arguments.
type Range [2]float64

// Result holds the result of integration.
type Result struct {
	Value    float64 // estimated integral
	AbsError float64 // estimated absolute error
	NEval    int     // number of function evaluations
}
