# Heuristic Optimization Algorithms
Provides Golang implementations of various heuristic optimization algorithms.<br>
The code is not meant to be production ready nor super optimized.<br>
Though it can be used for educational and research related tasks.

## Installation:
On your project's root folder (the one which contains go.mod) do run:
```
go get github.com/applied-math-coding/heuristic
```

## Example usage:

All examples require you have added the following dependencies:
```
import (
	"fmt"
	"math"	
	"gonum.org/v1/gonum/mat"
)
````

### Particle-Swarm optimization:
Add "github.com/applied-math-coding/heuristic/pso" to your imports.<br>
The following code searches a minimum of the function f:
```
func main() {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	params := &pso.Params{
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
	min := pso.Optimize(f, b_low, b_up, params)
	fmt.Println(min)
	fmt.Println(f(min))
}
```

### Artificial-Bees-Colony optimization:
Add "github.com/applied-math-coding/heuristic/abc" to your imports.<br>
The following code searches a minimum of the function f:

```
func main() {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	params := &abc.Params{
		N_bees:        100,
		Abandon_limit: 10,
		Max_iter:      100}
	f := func(x mat.Vector) float64 {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return math.Pow((1.0-x1)*x0, 2.0) + math.Pow(x1*(2.0-x0), 2.0)
	}
	min := abc.Optimize(f, b_low, b_up, params)
	fmt.Println(min)
	fmt.Println(f(min))
}
```

### Differential-Evolution optimization:
Add "github.com/applied-math-coding/heuristic/de" to your imports.<br>
The following code searches a minimum of the function f:
```
func main() {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	params := &de.Params{
		N_agents: 2000,
		Max_iter: 500,
		F:        0.8,
		CR:       0.9}
	f := func(x mat.Vector) float64 {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return math.Pow((1.0-x1)*x0, 2.0) + math.Pow(x1*(2.0-x0), 2.0)
	}
	min := de.Optimize(f, b_low, b_up, params)
	fmt.Println(min)
	fmt.Println(f(min))
}
```

### Meta-Optimizer (LUS):
Add "github.com/applied-math-coding/heuristic/lus" to your imports.<br>
The following code searches a minimum of the function f:
```
func main() {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	f := func(x mat.Vector) float64 {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return math.Pow((1.0-x1)*x0, 2.0) + math.Pow(x1*(2.0-x0), 2.0)
	}
	min := lus.Optimize(b_low, b_up, &lus.Params{Max_iter: 1000, Precision: 0.01}, f)
	fmt.Println(min)
	fmt.Println(f(min))
}
```
A use case of how this can be applied in order to supply PSO with optimal parameter, can be find at "./meta_opt_pso/meta_opt_pso.go".

### Global Root-Finder:
Add "github.com/applied-math-coding/heuristic/roots" to your imports.<br>
The following tries to find all roots for a function f which maps R^n to R^m. Internally it applies an interval-bisection
to search for local roots at various positions. The roots are search by applying the PSO onto |f|^2.
```
func main() {
	b_low := mat.NewVecDense(2, []float64{-10.0, -10.0})
	b_up := mat.NewVecDense(2, []float64{10.0, 10.0})
	f := func(x mat.Vector) mat.Vector {
		x0 := x.AtVec(0)
		x1 := x.AtVec(1)
		return mat.NewVecDense(2, []float64{
			(1.0 - x1) * math.Sin(x0),
			x1 * (2.0 - x0)})
	}
	params := &roots.Params{
		Location_Precision: 0.5,
		Root_Recognition:   0.1,
		N_particles:        500,
		Precision:          0.0000001}
	roots := roots.FindRoots(f, nil, b_low, b_up, params)
	for _, r := range roots {
		fmt.Println(r)
	}
}
```

## Further resources
A good overview about<br>

PSO:<br>
https://en.wikipedia.org/wiki/Particle_swarm_optimization

ABC:<br>
http://www.scholarpedia.org/article/Artificial_bee_colony_algorithm

DE:<br>
https://en.wikipedia.org/wiki/Differential_evolution

Meta-optimization (LSU):<br>
https://www.sciencedirect.com/science/article/abs/pii/S1568494609001549?casa_token=VeVpKW0fOlkAAAAA:i60-D9lgWAZnnklmihTMEjztvm3ZtGJVXdvCYIE9AUtSM4xsa3TzwCtRbbaLczDaJezsaOSIDr0



