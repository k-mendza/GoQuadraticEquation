package quadraticEquation

import (
	"sort"
	"testing"
)

func TestTableQuadraticEquation_GetRoots(t *testing.T) {
	var tests = []struct {
		input    *QuadraticEquation
		expected []float64
	}{
		{NewQuadraticEquation(1, 2, -3), []float64{-3, 1}},
		{NewQuadraticEquation(1, 2, 1), []float64{-1}},
		{NewQuadraticEquation(1, 1, 1), []float64{}},
	}
	for _, test := range tests {
		if output := test.input.GetRoots(); notEqual(output, test.expected) {
			t.Error("Test failed: {} input - {} expected - {} received", test.input, test.expected, output)
		}
	}

}

func notEqual(source []float64, target []float64) bool {
	if len(source) != len(target) {
		return true
	}
	sort.Float64s(source)
	sort.Float64s(target)
	for i, srcElement := range source {
		if srcElement != target[i] {
			return true
		}
	}
	return false
}
