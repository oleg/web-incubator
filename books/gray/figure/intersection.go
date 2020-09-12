package figure

import (
	"gray/oned"
	"sort"
)

type Inter struct {
	Distance float64
	Object   Shape
}

func (i Inter) PrepareComputations(r Ray) Computations {
	comps := Computations{}
	comps.Distance = i.Distance
	comps.Object = i.Object
	comps.Point = r.Position(comps.Distance)
	comps.EyeV = r.Direction.Negate()
	comps.NormalV = comps.Object.NormalAt(comps.Point)
	if comps.NormalV.Dot(comps.EyeV) < 0 {
		comps.Inside = true
		comps.NormalV = comps.NormalV.Negate()
	}
	return comps
}

type Computations struct {
	Distance float64
	Object   Shape
	Point    oned.Point
	EyeV     oned.Vector
	NormalV  oned.Vector
	Inside   bool
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
