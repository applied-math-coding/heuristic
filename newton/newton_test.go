package newton

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

var f = func(x mat.Vector) mat.Vector {
	x0 := x.AtVec(0)
	x1 := x.AtVec(1)
	return mat.NewVecDense(2, []float64{
		(1.0 - x1) * x0,
		x1 * (2.0 - x0)})
}

var x_0 = mat.NewVecDense(2, []float64{2.5, 1.5})

func TestFindRoot(t *testing.T) {
	D := func(x mat.Vector) mat.Matrix {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return mat.NewDense(2, 2, []float64{
			1.0 - x1, -x0,
			-x1, 2 - x0})
	}
	r, errFindRoot := FindRoot(
		f,
		D,
		x_0,
		&Params{Max_iter: 1000, Precision: 0.00001})
	if errFindRoot != nil {
		t.Fatal("TestFindRoot fails")
	} else {
		t.Log(r)
	}
}

func TestApproximateDerivative(t *testing.T) {
	D_x_0 := ApproximateDerivative(f, x_0)
	t.Logf("J ≈ %.6v\n", mat.Formatted(D_x_0, mat.Prefix("    ")))
}

func TestFindRootNoD(t *testing.T) {
	r, errFindRoot := FindRoot(
		f,
		nil,
		x_0,
		&Params{Max_iter: 1000, Precision: 0.00001})
	if errFindRoot != nil {
		t.Fatal("TestFindRoot fails")
	} else {
		t.Log(r)
	}
}
