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

	normalV := NormalAt(comps.Object, comps.Point)
	comps.Inside = normalV.Dot(comps.EyeV) < 0
	if comps.Inside {
		comps.NormalV = normalV.Negate()
	} else {
		comps.NormalV = normalV
	}
	comps.OverPoint = comps.Point.AddVector(comps.NormalV.MultiplyScalar(oned.Delta))
	comps.ReflectV = r.Direction.Reflect(comps.NormalV)
	return comps
}

type Computations struct {
	Distance  float64
	Object    Shape
	Point     oned.Point
	EyeV      oned.Vector
	NormalV   oned.Vector
	Inside    bool
	OverPoint oned.Point
	ReflectV  oned.Vector
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
