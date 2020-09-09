package figure

import "sort"

type Inter struct {
	Distance float64
	Object   Shape
}

type Inters []Inter

func (xs Inters) Hit() (bool, Inter) {
	//todo move to constructor
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].Distance < xs[j].Distance
	})
	for _, e := range xs {
		if e.Distance > 0 {
			return true, e
		}
	}
	return false, Inter{}
}
