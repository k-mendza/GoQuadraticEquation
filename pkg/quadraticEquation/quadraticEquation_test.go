package quadraticEquation

import (
	"testing"
)

func TestTableQuadraticEquation_GetRoots(t *testing.T) {
	var tests = []struct {
		input *QuadraticEquation
		expected []float64
	}{
		{NewQuadraticEquation(1, 2, -3), []float64{-3, 1}},
	}

	for _, test :=range tests {
		if output := test.input.GetRoots(); notEqual(output, test.expected) {
			t.Error("Test failed: {} input - {} expected - {} received", test.input, test.expected, output)
		}
	}

}

func notEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return true
		}
	}
	return false
}
