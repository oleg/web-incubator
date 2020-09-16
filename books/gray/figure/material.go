package figure

import (
	"gray/oned"
	"math"
)

//todo change types?
//todo reorder members
type Material struct {
	Color      oned.Color
	Ambient    float64
	Diffuse    float64
	Specular   float64
	Shininess  float64
	Pattern    Pattern
	Reflective float64
}

//todo change api, should accept overrides, use builder?
func DefaultMaterial() Material {
	return Material{
		Color:     oned.White,
		Ambient:   0.1,
		Specular:  0.9,
		Diffuse:   0.9,
		Shininess: 200.0,
	}
}

func Lighting(material Material, object Shape, light PointLight, point oned.Point, eyev oned.Vector, normalv oned.Vector, inShadow bool) oned.Color {
	var color oned.Color
	if material.Pattern != nil {
		color = PatternAtShape(material.Pattern, object, point)
	} else {
		color = material.Color
	}
	effectiveColor := color.Multiply(light.Intensity)
	lightv := light.Position.SubtractPoint(point).Normalize()
	ambient := effectiveColor.MultiplyByScalar(material.Ambient)
	lightDotNormal := lightv.Dot(normalv)
	if lightDotNormal < 0 || inShadow {
		return ambient
	}
	diffuse := effectiveColor.MultiplyByScalar(material.Diffuse).MultiplyByScalar(lightDotNormal)
	reflectv := lightv.Negate().Reflect(normalv)
	reflectDotEye := reflectv.Dot(eyev)
	if reflectDotEye <= 0 {
		return ambient.Add(diffuse)
	}
	factor := math.Pow(reflectDotEye, material.Shininess)
	specular := light.Intensity.MultiplyByScalar(material.Specular).MultiplyByScalar(factor)
	return ambient.Add(diffuse).Add(specular)
}

//todo think about it
type MaterialBuilder struct {
	color      oned.Color
	ambient    float64
	diffuse    float64
	specular   float64
	shininess  float64
	pattern    Pattern
	reflective float64
}

func MakeMaterialBuilder() *MaterialBuilder {
	return &MaterialBuilder{
		color:     oned.White,
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}
func (mb *MaterialBuilder) Build() Material {
	return Material{
		Color:      mb.color,
		Ambient:    mb.ambient,
		Diffuse:    mb.diffuse,
		Specular:   mb.specular,
		Shininess:  mb.shininess,
		Pattern:    mb.pattern,
		Reflective: mb.reflective,
	}
}
func (mb *MaterialBuilder) SetColor(color oned.Color) *MaterialBuilder {
	mb.color = color
	return mb
}
func (mb *MaterialBuilder) SetAmbient(ambient float64) *MaterialBuilder {
	mb.ambient = ambient
	return mb
}
func (mb *MaterialBuilder) SetDiffuse(diffuse float64) *MaterialBuilder {
	mb.diffuse = diffuse
	return mb
}
func (mb *MaterialBuilder) SetSpecular(specular float64) *MaterialBuilder {
	mb.specular = specular
	return mb
}
func (mb *MaterialBuilder) SetShininess(shininess float64) *MaterialBuilder {
	mb.shininess = shininess
	return mb
}
func (mb *MaterialBuilder) SetPattern(pattern Pattern) *MaterialBuilder {
	mb.pattern = pattern
	return mb
}
func (mb *MaterialBuilder) SetReflective(reflective float64) *MaterialBuilder {
	mb.reflective = reflective
	return mb
}
