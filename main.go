package main

import (
	"fmt"
	"main/roots"
	"math"

	"gonum.org/v1/gonum/mat"
)

func main() {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	f := func(x mat.Vector) mat.Vector {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return mat.NewVecDense(2, []float64{
			(1.0 - x1) * math.Sin(x0),
			x1 * (2.0 - x0)})
	}
	params := &roots.Params{
		Location_Precision: 0.5,
		Root_Recognition:   0.1,
		N_particles:        500,
		Precision:          0.0000001}
	roots := roots.FindRoots(f, nil, b_low, b_up, params)
	for _, r := range roots {
		fmt.Println(r)
	}
}
