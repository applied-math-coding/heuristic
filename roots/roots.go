package roots

import (
	"main/common"
	"main/meta_opt_pso"
	"main/newton"
	"main/pso"
	"math"

	"gonum.org/v1/gonum/mat"
)

type Params = struct {
	Root_Recognition   float64 // f(x) to be recognized as root by particles
	Location_Precision float64 // min distance to distinguish several roots
	N_particles        int     // based on the dimension and size of interval, one must play with this
	Precision          float64 // aimed precission of root
}

type Segment = struct {
	idx   int
	roots []mat.Vector
}

// FindRoots tries to find all roots of f in given multi-dim-interval. The algorithm uses a heuristic search (PSO)
// which imposes no restrictions on f. A recursive bisection procedure ensures the pso-search to provide
// as much as possible independent results.
// In case the system is n*n and a derivative is supplied, it is used for a Newton-method to refine roots.
// If the system is n*n and no derivative is supplied, the derivative will be approximated internally.
func FindRoots(f common.System, D common.Derivative, b_low mat.Vector, b_up mat.Vector,
	params *Params) []mat.Vector {
	pso_params := meta_opt_pso.Optimize(createTargetFn(f), b_low, b_up)
	pso_params.Max_iter = 100
	pso_params.N_particles = int(math.Max(500.0, float64(params.N_particles)))
	roots := recFindRoots(f, b_low, b_up, params, pso_params, nil)
	res := make([]mat.Vector, 0)
	for _, r := range roots {
		if !isRootContainedInList(params.Location_Precision, r, res) {
			res = append(res, r)
		}
	}
	m := f(b_low).Len()
	if m == b_low.Len() {
		return RefineRoots(f, D, res, params)
	}
	return res
}

func RefineRoots(f common.System, D common.Derivative, roots []mat.Vector, params *Params) []mat.Vector {
	res := make([]mat.Vector, len(roots))
	for idx, r := range roots {
		refined, errFindRoot := newton.FindRoot(
			f, D, r, &newton.Params{Max_iter: 1000, Precision: params.Precision})
		if errFindRoot != nil {
			res[idx] = r
		} else {
			res[idx] = refined
		}
	}
	return res
}

func isRootContainedInList(location_Precision float64, r mat.Vector, list []mat.Vector) bool {
	for _, e := range list {
		isSame := true
		for idx := 0; idx < r.Len(); idx++ {
			if math.Abs(r.AtVec(idx)-e.AtVec(idx)) > location_Precision {
				isSame = false
				break
			}
		}
		if isSame {
			return true
		}
	}
	return false
}

func isUnderPrecision(b_low mat.Vector, b_up mat.Vector, precision float64) bool {
	for idx := 0; idx < b_low.Len(); idx++ {
		if math.Abs(b_low.AtVec(idx)-b_up.AtVec(idx)) < precision {
			return true
		}
	}
	return false
}

func recFindRoots(f common.System, b_low mat.Vector, b_up mat.Vector,
	params *Params, pso_params *pso.Params, segment *Segment) []mat.Vector {
	res := make([]mat.Vector, 0)
	if segment == nil {
		segment = &Segment{idx: 0, roots: make([]mat.Vector, 0)}
	}
	if isUnderPrecision(b_low, b_up, params.Location_Precision) {
		return res
	}
	if len(segment.roots) > 0 {
		res = append(res, findRootsInDeeperLevels(f, b_low, b_up, params, pso_params, segment)...)
	} else {
		root := searchRoot(f, b_low, b_up, pso_params, params)
		if root == nil {
			return res
		} else {
			segment.roots = append(segment.roots, root)
			res = append(res, root)
			res = append(res, findRootsInDeeperLevels(f, b_low, b_up, params, pso_params, segment)...)
		}
	}
	return res
}

func findRootsInDeeperLevels(f common.System, b_low mat.Vector, b_up mat.Vector,
	params *Params, pso_params *pso.Params, segment *Segment) []mat.Vector {
	res := make([]mat.Vector, 0)
	idx := int(math.Mod(float64(segment.idx)+1.0, float64(b_low.Len())))
	b_center_up, b_center_low := splitInterval(idx, b_low, b_up)
	segment_low_roots := make([]mat.Vector, 0)
	segment_up_roots := make([]mat.Vector, 0)
	for _, r := range segment.roots {
		if isContained(r, b_low, b_center_up) {
			segment_low_roots = append(segment_low_roots, r)
		} else {
			segment_up_roots = append(segment_up_roots, r)
		}
	}
	segment_low := &Segment{idx: idx, roots: segment_low_roots}
	res = append(res, recFindRoots(f, b_low, b_center_up, params, pso_params, segment_low)...)
	segment_up := &Segment{idx: idx, roots: segment_up_roots}
	res = append(res, recFindRoots(f, b_center_low, b_up, params, pso_params, segment_up)...)
	return res
}

func isContained(v mat.Vector, b_low mat.Vector, b_up mat.Vector) bool {
	res := true
	for idx := 0; idx < v.Len(); idx++ {
		res = v.AtVec(idx) >= b_low.AtVec(idx) && v.AtVec(idx) < b_up.AtVec(idx)
	}
	return res
}

func createTargetFn(f common.System) common.Target {
	return func(x mat.Vector) float64 {
		res := 0.0
		y := f(x)
		for i := 0; i < y.Len(); i++ {
			res = res + math.Pow(y.AtVec(i), 2.0)
		}
		return res
	}
}

func searchRoot(f common.System, b_low mat.Vector, b_up mat.Vector, pso_params *pso.Params, params *Params) mat.Vector {
	x_0 := pso.Optimize(createTargetFn(f), b_low, b_up, pso_params)
	isRoot := true
	y_0 := f(x_0)
	for i := 0; i < x_0.Len(); i++ {
		isRoot = math.Abs(y_0.AtVec(i)) <= params.Root_Recognition
		if !isRoot {
			return nil
		}
	}
	return x_0
}

func splitInterval(idx int, b_low mat.Vector, b_up mat.Vector) (mat.Vector, mat.Vector) {
	n := b_low.Len()
	mid := b_low.AtVec(idx) + 0.5*(b_up.AtVec(idx)-b_low.AtVec(idx))
	b_center_up := mat.NewVecDense(n, nil)
	b_center_up.CopyVec(b_up)
	b_center_up.SetVec(idx, mid)
	b_center_low := mat.NewVecDense(n, nil)
	b_center_low.CopyVec(b_low)
	b_center_low.SetVec(idx, mid)
	return b_center_up, b_center_low
}
