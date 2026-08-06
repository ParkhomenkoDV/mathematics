// Package integrate provides numerical integration routines.
package integrate

const (
	EpsAbs = 1.49e-8
	EpsRel = 1.49e-8
	Limit  = 50
)

type Func func(...float64) float64

type Range [2]float64 // TODO: func

type Result struct {
	Value    float64 // estimated integral
	AbsError float64 // estimated absolute error
	NEval    int     // number of function evaluations
}
