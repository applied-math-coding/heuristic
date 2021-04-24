package meta_opt_pso

import (
	"main/common"
	"main/lus"
	"main/pso"

	"gonum.org/v1/gonum/mat"
)

func Optimize(f common.Target, b_low mat.Vector, b_up mat.Vector) *pso.Params {
	phi_max := 4.0
	omega_max := 2.0
	learning_rate_max := 1.0
	lus_b_low := mat.NewVecDense(4, []float64{-omega_max, -phi_max, -phi_max, 0.01})
	lus_b_up := mat.NewVecDense(4, []float64{omega_max, phi_max, phi_max, learning_rate_max})
	optimal := lus.Optimize(
		lus_b_low,
		lus_b_up,
		&lus.Params{Max_iter: 1000, Precision: 0.01},
		func(x mat.Vector) float64 {
			return f(pso.Optimize(f, b_low, b_up, &pso.Params{
				Omega:        x.AtVec(0),
				Phi_p:        x.AtVec(1),
				Phi_g:        x.AtVec(2),
				N_particles:  1000,
				LearningRate: x.AtVec(3),
				Max_iter:     10}))
		})
	return &pso.Params{
		Omega:        optimal.AtVec(0),
		Phi_p:        optimal.AtVec(1),
		Phi_g:        optimal.AtVec(2),
		LearningRate: optimal.AtVec(3)}
}
