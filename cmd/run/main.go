package main

import (
	"fmt"
	qe "quadraticEquation/pkg/quadraticEquation"
)

func main() {
	var quadEq = qe.NewQuadraticEquation(1,2,-3)
	fmt.Println(quadEq.GetRoots())
}
