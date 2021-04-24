package pso

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestPso(t *testing.T) {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	params := &Params{
		N_particles:  500,
		Max_iter:     100,
		Omega:        1.0,
		Phi_p:        2.0,
		Phi_g:        2.0,
		LearningRate: 0.5}
	f := func(x mat.Vector) float64 {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return math.Pow((1.0-x1)*x0, 2.0) + math.Pow(x1*(2.0-x0), 2.0)
	}
	min := Optimize(f, b_low, b_up, params)
	t.Log(min)
	t.Log(f(min))
}

func TestInitVelocity(t *testing.T) {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	v := initVelocity(b_low, b_up, 3)
	t.Log(v[0])
	t.Log(v[1])
	t.Log(v[2])
}

func TestInitParticles(t *testing.T) {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	p := initParticles(b_low, b_up, 3)
	t.Log(p[0])
	t.Log(p[1])
	t.Log(p[2])
}

func TestParticlePositions(t *testing.T) {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	p := initParticles(b_low, b_up, 3)
	x := initParticlePositions(p)
	t.Log(x[0])
	t.Log(x[1])
	t.Log(x[2])
}
