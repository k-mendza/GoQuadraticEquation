package quadraticEquation

import "math"

type QuadraticEquation struct {
	secondPowerCoeff float64
	firstPowerCoeff  float64
	zeroPowerCoeff   float64
	delta            float64
}

func NewQuadraticEquation(secondPowerCoeff float64, firstPowerCoeff float64, zeroPowerCoeff float64) *QuadraticEquation {
	equation := new(QuadraticEquation)
	equation.secondPowerCoeff = secondPowerCoeff
	equation.firstPowerCoeff = firstPowerCoeff
	equation.zeroPowerCoeff = zeroPowerCoeff
	equation.delta = math.Pow(firstPowerCoeff, 2) - 4*secondPowerCoeff*zeroPowerCoeff
	return equation
}

func (eq *QuadraticEquation) GetRoots() []float64 {
	if eq.isDeltaZero() {
		return eq.getZeroDeltaRoot()
	}

	if eq.isDeltaPositive() {
		return eq.getPositiveDeltaRoots()
	}

	return []float64{}
}

func (eq *QuadraticEquation) isDeltaPositive() bool {
	return eq.delta >= 0
}

func (eq *QuadraticEquation) isDeltaZero() bool {
	return eq.delta == 0
}

func (eq *QuadraticEquation) getZeroDeltaRoot() []float64 {
	return []float64{-eq.firstPowerCoeff / (2 * eq.secondPowerCoeff)}
}

func (eq *QuadraticEquation) getPositiveDeltaRoots() []float64 {
	return []float64{(-eq.firstPowerCoeff - math.Sqrt(eq.delta)) / (2 * eq.secondPowerCoeff),
		(-eq.firstPowerCoeff + math.Sqrt(eq.delta)) / (2 * eq.secondPowerCoeff)}
}
