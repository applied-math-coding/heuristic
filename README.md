# Heuristic Optimization Algorithms
Golang implementation of various heuristic optimization algorithms.
The code is not meant to be production ready nor super optimized. 
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


