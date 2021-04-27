package newton

import (
	"main/common"
	"math"

	"gonum.org/v1/gonum/diff/fd"
	"gonum.org/v1/gonum/mat"
)

type Params = struct {
	Max_iter  int
	Precision float64
}

// FindRoot uses the simplified Newton method in order to approximate a root.
// The derivate is required to be invertable at x_0, otherwise error is thrown.
// D is optional and if not given it will be approximated.
func FindRoot(f common.System, D common.Derivative,
	x_0 mat.Vector, params *Params) (mat.Vector, error) {
	x := mat.NewVecDense(x_0.Len(), nil)
	x.CopyVec(x_0)
	D_inv := mat.NewDense(x_0.Len(), x_0.Len(), nil)
	var D_x0 mat.Matrix
	if D != nil {
		D_x0 = D(x_0)
	} else {
		D_x0 = ApproximateDerivative(f, x_0)
	}
	if err_inverse := D_inv.Inverse(D_x0); err_inverse != nil {
		return nil, err_inverse
	}
	value := mat.Norm(f(x_0), 2)
	for iter := 0; math.Abs(value) > float64(params.Precision) && iter <= params.Max_iter; iter++ {
		z := mat.NewVecDense(x.Len(), nil)
		z.MulVec(D_inv, f(x))
		x.SubVec(x, z)
		value = mat.Norm(f(x), 2)
	}
	return x, nil
}

func ApproximateDerivative(f common.System, x_0 mat.Vector) mat.Matrix {
	ff := func(yy, x []float64) {
		y := f(mat.NewVecDense(x_0.Len(), x))
		for i := 0; i < x_0.Len(); i++ {
			yy[i] = y.AtVec(i)
		}
	}
	dst := mat.NewDense(x_0.Len(), x_0.Len(), nil)
	xx_0 := x_0.(*mat.VecDense).RawVector().Data
	fd.Jacobian(dst, ff, xx_0, nil)
	return dst
}
