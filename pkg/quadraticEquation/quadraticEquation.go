package quadraticEquation

import "math"

type QuadraticEquation struct {
	secondPowerCoeff float64
	firstPowerCoeff float64
	zeroPowerCoeff float64
}

func NewQuadraticEquation(secondPowerCoeff float64, firstPowerCoeff float64, zeroPowerCoeff float64) *QuadraticEquation {
	equation := new(QuadraticEquation)
	equation.secondPowerCoeff = secondPowerCoeff
	equation.firstPowerCoeff = firstPowerCoeff
	equation.zeroPowerCoeff = zeroPowerCoeff
	return equation
}

func (eq *QuadraticEquation) GetRoots() []float64 {
	if eq.isDeltaPositive() {
		if eq.isDeltaZero() {
			return eq.getZeroDeltaRoot()
		}
		return eq.getPositiveDeltaRoots()
	}
	return []float64{}
}

func (eq *QuadraticEquation) isDeltaPositive() bool {
	return eq.delta() >= 0
}

func (eq *QuadraticEquation) isDeltaZero() bool {
	return eq.delta() == 0
}

func (eq *QuadraticEquation) delta() float64 {
	return math.Pow(eq.firstPowerCoeff, 2)-4*eq.secondPowerCoeff*eq.zeroPowerCoeff
}

func (eq *QuadraticEquation) getZeroDeltaRoot() []float64 {
	return []float64{-eq.firstPowerCoeff/(2*eq.secondPowerCoeff)}
}

func (eq *QuadraticEquation) getPositiveDeltaRoots() []float64 {
	var deltaSquared = math.Sqrt(eq.delta())
	var twoTimesSecondPowerCoeff = 2*eq.secondPowerCoeff
	
	return []float64{(-eq.firstPowerCoeff-deltaSquared)/(twoTimesSecondPowerCoeff),
		(-eq.firstPowerCoeff+deltaSquared)/(twoTimesSecondPowerCoeff)}
}