package figure

import (
	"github.com/stretchr/testify/assert"
	"gray/multid"
	"gray/oned"
	"math"
	"testing"
)

func Test_default_world(t *testing.T) {
	light := PointLight{oned.Point{-10, 10, -10}, oned.White}

	material := DefaultMaterial()
	material.Color = oned.Color{0.8, 1.0, 0.6}
	material.Diffuse = 0.7
	material.Specular = 0.2
	s1 := MakeSphereM(material)

	transform := multid.Scaling(0.5, 0.5, 0.5)
	s2 := MakeSphereT(transform)

	w := defaultWorld()

	assert.Equal(t, light, w.Light)
	assert.Equal(t, s1, w.Objects[0])
	assert.Equal(t, s2, w.Objects[1])
}

func Test_Intersect_world_with_ray(t *testing.T) {
	w := defaultWorld()
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}

	xs := w.Intersect(r)

	assert.Equal(t, 4, len(xs))
	assert.Equal(t, 4.0, xs[0].Distance)
	assert.Equal(t, 4.5, xs[1].Distance)
	assert.Equal(t, 5.5, xs[2].Distance)
	assert.Equal(t, 6.0, xs[3].Distance)
}

func Test_shading_intersection(t *testing.T) {
	w := defaultWorld()
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}
	shape := w.Objects[0]
	i := Inter{4, shape}
	comps := i.prepareComputations(r)

	c := w.ShadeHit(comps, MaxDepth)

	oned.AssertColorEqualInDelta(t, oned.Color{0.38066, 0.47583, 0.2855}, c)
}

func Test_shading_intersection_from_inside(t *testing.T) {
	w := defaultWorld()
	w.Light = PointLight{oned.Point{0, 0.25, 0}, oned.White}
	r := Ray{oned.Point{0, 0, 0}, oned.Vector{0, 0, 1}}
	shape := w.Objects[1]
	i := Inter{0.5, shape}
	comps := i.prepareComputations(r)

	c := w.ShadeHit(comps, MaxDepth)

	oned.AssertColorEqualInDelta(t, oned.Color{0.90498, 0.90498, 0.90498}, c)
}

func Test_color_when_ray_misses(t *testing.T) {
	w := defaultWorld()
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 1, 0}}

	c := w.ColorAt(r, MaxDepth)

	oned.AssertColorEqualInDelta(t, oned.Black, c)
}

func Test_color_when_ray_hits(t *testing.T) {
	w := defaultWorld()
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}

	c := w.ColorAt(r, MaxDepth)

	oned.AssertColorEqualInDelta(t, oned.Color{0.38066, 0.47583, 0.2855}, c)
}

func Test_color_with_intersection_behind_ray(t *testing.T) {
	w := World{pointLightSample(), []Shape{
		MakeSphereM(testMaterialBuilder().SetAmbient(1).Build()),
		MakeSphereTM(multid.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())}}
	r := Ray{oned.Point{0, 0, 0.75}, oned.Vector{0, 0, -1}}

	c := w.ColorAt(r, MaxDepth)

	oned.AssertColorEqualInDelta(t, testMaterialBuilder().SetAmbient(1).Build().Color, c)
}

func Test_shade_hit_is_given_intersection_in_shadow(t *testing.T) {
	s1 := MakeSphere()
	s2 := MakeSphereT(multid.Translation(0, 0, 10))
	w := World{
		PointLight{oned.Point{0, 0, -10}, oned.White},
		[]Shape{s1, s2},
	}
	r := Ray{oned.Point{0, 0, 5}, oned.Vector{0, 0, 1}}
	i := Inter{4, s2}
	comps := i.prepareComputations(r)

	color := w.ShadeHit(comps, MaxDepth)

	assert.Equal(t, oned.Color{0.1, 0.1, 0.1}, color)
}

func Test_hit_should_offset_point(t *testing.T) {
	r := Ray{oned.Point{0, 0, -5}, oned.Vector{0, 0, 1}}
	s := MakeSphereT(multid.Translation(0, 0, 1))
	i := Inter{5, s}

	comps := i.prepareComputations(r)

	assert.Less(t, comps.OverPoint.Z, -oned.Delta/2)
	assert.Less(t, comps.OverPoint.Z, comps.Point.Z)
}

func Test_reflected_color_for_non_reflective_material(t *testing.T) {
	s1 := MakeSphereM(testMaterialBuilder().Build())
	s2 := MakeSphereTM(multid.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	w := World{pointLightSample(), []Shape{s1, s2}}
	r := Ray{oned.Point{0, 0, 0}, oned.Vector{0, 0, 1}}
	i := Inter{1, s2}
	comps := i.prepareComputations(r)

	color := w.ReflectedColor(comps, 5)

	assert.Equal(t, oned.Color{0, 0, 0}, color)
}

func Test_reflected_color_for_reflective_material(t *testing.T) {
	s1 := MakeSphereM(testMaterialBuilder().Build())
	s2 := MakeSphereTM(multid.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := MakePlaneTM(multid.Translation(0, -1, 0), MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []Shape{s1, s2, s3}}
	r := Ray{oned.Point{0, 0, -3}, oned.Vector{0, -math.Sqrt2 / 2, math.Sqrt2 / 2}}
	i := Inter{math.Sqrt2, s3}
	comps := i.prepareComputations(r)

	color := w.ReflectedColor(comps, 5)

	oned.AssertColorEqualInDelta(t, oned.Color{0.19033, 0.23791, 0.142749}, color)
}

func Test_shade_hit_with_reflective_material(t *testing.T) {
	s1 := MakeSphereM(testMaterialBuilder().Build())
	s2 := MakeSphereTM(multid.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := MakePlaneTM(multid.Translation(0, -1, 0), MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []Shape{s1, s2, s3}}
	r := Ray{oned.Point{0, 0, -3}, oned.Vector{0, -math.Sqrt2 / 2, math.Sqrt2 / 2}}
	i := Inter{math.Sqrt2, s3}
	comps := i.prepareComputations(r)

	color := w.ShadeHit(comps, MaxDepth)

	oned.AssertColorEqualInDelta(t, oned.Color{0.87675, 0.92434, 0.82918}, color)
}

func Test_color_at_with_mutually_reflective_surfaces(t *testing.T) {
	w := World{
		PointLight{oned.Point{0, 0, 0}, oned.Color{1, 1, 1}},
		[]Shape{
			MakePlaneTM(multid.Translation(0, -1, 0), MakeMaterialBuilder().SetReflective(1).Build()),
			MakePlaneTM(multid.Translation(0, 1, 0), MakeMaterialBuilder().SetReflective(1).Build())}}

	w.ColorAt(Ray{oned.Point{0, 0, 0}, oned.Vector{0, 1, 0}}, MaxDepth)

	//should terminate
}

func Test_reflected_color_at_maximum_recursive_depth(t *testing.T) {
	s1 := MakeSphereM(testMaterialBuilder().Build())
	s2 := MakeSphereTM(multid.Scaling(0.5, 0.5, 0.5), testMaterialBuilder().SetAmbient(1).Build())
	s3 := MakePlaneTM(multid.Translation(0, -1, 0), MakeMaterialBuilder().SetReflective(0.5).Build())
	w := World{pointLightSample(), []Shape{s1, s2, s3}}
	r := Ray{oned.Point{0, 0, -3}, oned.Vector{0, -math.Sqrt2 / 2, math.Sqrt2 / 2}}
	i := Inter{math.Sqrt2, s3}
	comps := i.prepareComputations(r)

	color := w.ReflectedColor(comps, 0)

	oned.AssertColorEqualInDelta(t, oned.Color{0, 0, 0}, color)
}

//util
func defaultWorld() World {
	s1 := MakeSphereM(testMaterialBuilder().Build())
	s2 := MakeSphereT(multid.Scaling(0.5, 0.5, 0.5))

	return World{
		pointLightSample(),
		[]Shape{s1, s2},
	}
}

func pointLightSample() PointLight {
	return PointLight{oned.Point{-10, 10, -10}, oned.White}
}

func testMaterialBuilder() *MaterialBuilder {
	return MakeMaterialBuilder().
		SetColor(oned.Color{0.8, 1.0, 0.6}).
		SetDiffuse(0.7).
		SetSpecular(0.2)
}
