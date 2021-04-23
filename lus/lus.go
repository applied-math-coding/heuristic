package lus

import (
	"main/common"
	"math"

	"gonum.org/v1/gonum/mat"
)

type Params = struct {
	Max_iter  int
	Precision float64
}

func Optimize(b_low mat.Vector, b_up mat.Vector, params *Params, f common.Target) mat.Vector {
	n := b_low.Len()
	beta := 1.0 / 3.0
	q := math.Pow(2.0, -beta/float64(n))
	x := common.RandomDataInBounds(b_low, b_up)
	d := mat.NewVecDense(n, nil)
	d.SubVec(b_up, b_low)
	best_value := f(x)
	diam := math.Max(math.Abs(mat.Max(b_low)), math.Abs(mat.Max(b_up)))
	for iter := 0; iter < params.Max_iter && diam > params.Precision; iter++ {
		d_m := mat.NewVecDense(n, nil)
		d_m.ScaleVec(-1.0, d)
		a := common.RandomDataInBounds(d_m, d)
		y := mat.NewVecDense(n, nil)
		y.AddVec(x, a)
		value := f(y)
		if value < best_value {
			x = y
			if math.Abs(best_value-value) < params.Precision {
				break
			}
			best_value = value
		} else {
			d.ScaleVec(q, d)
			diam = q * diam
		}
	}
	return x
}
