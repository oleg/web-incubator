package figure

import (
	"gray/oned"
	"sort"
)

type World struct {
	Light   PointLight
	Objects []Shape
}

func (w World) Intersect(ray Ray) Inters {
	r := Inters{}
	for _, shape := range w.Objects {
		r = append(r, shape.Intersect(ray)...)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].Distance < r[j].Distance
	})
	return r
}

func (w World) ShadeHit(comps Computations) oned.Color {
	return Lighting(comps.Object.Material(), w.Light, comps.Point, comps.EyeV, comps.NormalV)
}

func (w World) ColorAt(r Ray) oned.Color {
	if hit, inter := w.Intersect(r).Hit(); hit {
		return w.ShadeHit(inter.PrepareComputations(r))
	}
	return oned.Color{}
}
