package meta_opt_pso

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestOptimize(t *testing.T) {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	f := func(x mat.Vector) float64 {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return math.Pow((1.0-x1)*x0, 2.0) + math.Pow(x1*(2.0-x0), 2.0)
	}
	t.Log(Optimize(f, b_low, b_up))
}
