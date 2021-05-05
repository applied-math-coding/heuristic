package pso

import (
	"math"

	"github.com/applied-math-coding/heuristic/common"

	"gonum.org/v1/gonum/mat"
)

type Params = struct {
	Omega        float64
	Phi_p        float64
	Phi_g        float64
	N_particles  int
	LearningRate float64
	Max_iter     int
}

func Optimize(f common.Target, b_low mat.Vector, b_up mat.Vector, params *Params) mat.Vector {
	g := common.RandomDataInBounds(b_low, b_up)
	p := initParticles(b_low, b_up, params.N_particles)
	v := initVelocity(b_low, b_up, params.N_particles)
	x := initParticlePositions(p)
	value_g := f(g)
	for iter := 0; iter < params.Max_iter; iter++ {
		for i := 0; i < params.N_particles; i++ {
			updateVelocity(v[i], x[i], p[i], g, params)
			updateParticlePositions(x[i], v[i], params.LearningRate, b_low, b_up)
			value_x := f(x[i])
			if value_x < f(p[i]) {
				p[i].CopyVec(x[i])
			}
			if value_x < value_g {
				value_g = value_x
				g.CopyVec(x[i])
			}
		}
	}
	return g
}

func updateParticlePositions(x *mat.VecDense, v *mat.VecDense, learningRate float64,
	b_low mat.Vector, b_up mat.Vector) {
	x.AddScaledVec(x, learningRate, v)
	common.ApplyElementwise(x, func(e float64, idx int) float64 {
		return math.Min(math.Max(b_low.AtVec(idx), e), b_up.AtVec(idx))
	})
}

func updateVelocity(v *mat.VecDense, x *mat.VecDense,
	p *mat.VecDense, g *mat.VecDense, params *Params) {
	n := g.Len()
	r := common.GetNewRand()
	r_p := r.Float64()
	r_g := r.Float64()
	v.ScaleVec(params.Omega, v)
	diff_p_x := mat.NewVecDense(n, nil)
	diff_p_x.SubVec(p, x)
	diff_p_x.ScaleVec(r_p*params.Phi_p, diff_p_x)
	v.AddVec(v, diff_p_x)
	diff_g_x := mat.NewVecDense(n, nil)
	diff_g_x.SubVec(g, x)
	diff_g_x.ScaleVec(r_g*params.Phi_g, diff_g_x)
	v.AddVec(v, diff_g_x)
}

func initParticlePositions(p []*mat.VecDense) []*mat.VecDense {
	n := p[0].Len()
	x := make([]*mat.VecDense, len(p))
	for i, e := range p {
		x[i] = mat.NewVecDense(n, nil)
		x[i].CopyVec(e)
	}
	return x
}

func initVelocity(b_low mat.Vector, b_up mat.Vector, n_particles int) []*mat.VecDense {
	n := b_low.Len()
	v := make([]*mat.VecDense, n_particles)
	diff := mat.NewVecDense(n, nil)
	diff.SubVec(b_up, b_low)
	common.ApplyElementwise(diff, func(e float64, idx int) float64 {
		return math.Abs(e)
	})
	lower := mat.NewVecDense(n, nil)
	lower.CopyVec(diff)
	lower.ScaleVec(-1.0, lower)
	for i := 0; i < n_particles; i++ {
		v[i] = common.RandomDataInBounds(lower, diff)
	}
	return v
}

func initParticles(b_low mat.Vector, b_up mat.Vector, n_particles int) []*mat.VecDense {
	p := make([]*mat.VecDense, n_particles)
	for i := 0; i < n_particles; i++ {
		p[i] = common.RandomDataInBounds(b_low, b_up)
	}
	return p
}
