package abc

import (
	"github.com/applied-math-coding/heuristic/common"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

type Params = struct {
	N_bees        int
	Abandon_limit int
	Max_iter      int
}

type BeeType = struct {
	position           mat.Vector
	value              float64
	not_improved_since int
	fitness            float64 // distance to worst
}

type Bee = *BeeType

func Optimize(f common.Target, b_low mat.Vector, b_up mat.Vector, params *Params) mat.Vector {
	employedBees := initBees(f, b_low, b_up, params.N_bees)
	best_bee := findBestBee(employedBees)
	best_position := &best_bee.position
	best_value := &best_bee.value
	for iter := 0; iter < params.Max_iter; iter++ {
		doEmployedBeesPhase(f, employedBees, best_position, best_value)
		doOnLookingPhase(employedBees)
		doScoutPhase(employedBees, params.Abandon_limit, b_low, b_up, f)
	}
	return *best_position
}

// abandon bees from position if have not been improved for long
func doScoutPhase(bees []Bee, abandon_limit int, b_low mat.Vector, b_up mat.Vector, f common.Target) {
	for _, b := range bees {
		if b.not_improved_since > abandon_limit {
			b.position = common.RandomDataInBounds(b_low, b_up)
			b.value = f(b.position)
			b.not_improved_since = 0
		}
	}
}

// re-distribute bees onto other locations and prioritize locations with higher fitness
func doOnLookingPhase(bees []Bee) {
	computeFitnessOnBees(bees)
	sumFitness := 0.0
	for _, b := range bees {
		sumFitness = sumFitness + b.fitness
	}
	positionWeights := make([]float64, len(bees))
	for i, b := range bees {
		positionWeights[i] = b.fitness / sumFitness
	}
	catDist := distuv.NewCategorical(positionWeights, rand.NewSource(uint64(common.GetNextSeed())))
	for _, b := range bees {
		i := int(catDist.Rand())
		b.position = bees[i].position
		b.value = bees[i].value
	}
}

func computeFitnessOnBees(bees []Bee) {
	_, worst := findValueBounds(bees)
	for _, b := range bees {
		b.fitness = worst - b.value
	}
}

func findValueBounds(bees []Bee) (float64, float64) {
	best := bees[0].value
	worst := bees[0].value
	for _, b := range bees {
		if b.value < best {
			best = b.value
		}
		if b.value > worst {
			worst = b.value
		}
	}
	return best, worst
}

// search in local neighborhood for better value
func doEmployedBeesPhase(f common.Target, bees []Bee, best_position *mat.Vector, best_value *float64) {
	ran := common.GetNewRand()
	for _, b := range bees {
		k := ran.Intn(len(bees))
		i := ran.Intn(b.position.Len())
		phi := -1.0 + 2.0*ran.Float64()
		v := mat.NewVecDense(b.position.Len(), nil)
		v.CopyVec(b.position)
		x_i := b.position.AtVec(i)
		x_k := bees[k].position.AtVec(i)
		v.SetVec(i, x_i+phi*(x_i-x_k))
		value := f(v)
		if value < b.value {
			b.position = v
			b.value = value
			b.not_improved_since = 0
		} else {
			b.not_improved_since = b.not_improved_since + 1
		}
		if b.value < *best_value {
			*best_value = b.value
			*best_position = b.position
		}
	}
}

func findBestBee(bees []Bee) Bee {
	best := bees[0]
	for _, b := range bees {
		if b.value < best.value {
			best = b
		}
	}
	return best
}

func initBees(f common.Target, b_low mat.Vector, b_up mat.Vector, n_bees int) []Bee {
	res := make([]Bee, n_bees)
	for i := 0; i < n_bees; i++ {
		position := common.RandomDataInBounds(b_low, b_up)
		res[i] = &BeeType{
			position: position,
			value:    f(position)}
	}
	return res
}
