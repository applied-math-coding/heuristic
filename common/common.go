package common

import (
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

type Target = func(x mat.Vector) float64
type System = func(x mat.Vector) mat.Vector
type Derivative = func(x mat.Vector) mat.Matrix

var EPSILON = math.Nextafter(1.0, 2.0) - 1.0

var seed int64 = 0

func ApplyElementwise(v *mat.VecDense, fn func(e float64, idx int) float64) {
	for i := 0; i < v.Len(); i++ {
		v.SetVec(i, fn(v.AtVec(i), i))
	}
}

func RandomDataInBounds(b_low mat.Vector, b_up mat.Vector) *mat.VecDense {
	r := GetNewRand()
	n := b_low.Len()
	diff := mat.NewVecDense(n, nil)
	diff.SubVec(b_up, b_low)
	data := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = r.Float64()
	}
	v := mat.NewVecDense(n, data)
	v.MulElemVec(v, diff)
	v.AddVec(b_low, v)
	return v
}

func GetNextSeed() int64 {
	seed++
	return seed
}

func GetRandomSource() int64 {
	return time.Now().UnixNano() + GetNextSeed()
}

func GetNewRand() *rand.Rand {
	return rand.New(rand.NewSource(GetRandomSource()))
}
