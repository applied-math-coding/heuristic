package de

import (
	"github.com/applied-math-coding/heuristic/common"

	"gonum.org/v1/gonum/mat"
)

type Params = struct {
	N_agents int
	F        float64
	CR       float64
	Max_iter int
}

func Optimize(f common.Target, b_low mat.Vector, b_up mat.Vector, params *Params) mat.Vector {
	agents := initAgents(b_low, b_up, params.N_agents)
	n := b_low.Len()
	global_best := findBest(f, agents)
	global_best_val := f(global_best)
	for iter := 0; iter <= params.Max_iter; iter++ {
		rand := common.GetNewRand()
		for agentIdx, x := range agents {
			a := agents[rand.Intn(n)]
			b := agents[rand.Intn(n)]
			c := agents[rand.Intn(n)]
			R := rand.Intn(n)
			y := mat.NewVecDense(n, nil)
			for i := 0; i < n; i++ {
				r := rand.Float64()
				if r < params.CR || i == R {
					y.SetVec(i, a.AtVec(i)+params.F*(b.AtVec(i)-c.AtVec(i)))
				} else {
					y.SetVec(i, x.AtVec(i))
				}
			}
			y_val := f(y)
			if y_val < f(x) {
				agents[agentIdx] = y
				if y_val < global_best_val {
					global_best = y
					global_best_val = y_val
				}
			}
		}
	}
	return global_best
}

func findBest(f common.Target, agents []mat.Vector) mat.Vector {
	best_agent := agents[0]
	best_val := f(best_agent)
	for _, x := range agents {
		x_val := f(x)
		if x_val < best_val {
			best_agent = x
			best_val = x_val
		}
	}
	return best_agent
}

func initAgents(b_low mat.Vector, b_up mat.Vector, n_agents int) []mat.Vector {
	res := make([]mat.Vector, n_agents)
	for i := 0; i < n_agents; i++ {
		res[i] = common.RandomDataInBounds(b_low, b_up)
	}
	return res
}
