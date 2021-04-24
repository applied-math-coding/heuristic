package main

import (
	"fmt"
	"main/roots"

	"gonum.org/v1/gonum/mat"
)

func main() {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	f := func(x mat.Vector) mat.Vector {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return mat.NewVecDense(2, []float64{
			(1.0 - x1) * x0,
			x1 * (2.0 - x0)})
	}
	params := &roots.Params{Max_level: 6, Precision: 0.1}
	roots := roots.FindRoots(f, b_low, b_up, params, nil, nil)
	for _, r := range roots {
		fmt.Println(r)
	}
}
