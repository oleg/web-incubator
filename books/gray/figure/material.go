package figure

import (
	"gray/oned"
	"math"
)

type Material struct {
	Color     oned.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

var DefaultMaterial = Material{
	Color:     oned.Color{1, 1, 1},
	Ambient:   0.1,
	Specular:  0.9,
	Diffuse:   0.9,
	Shininess: 200.0,
}

func Lighting(material Material, light PointLight, point oned.Point, eyev oned.Vector, normalv oned.Vector) oned.Color {
	effectiveColor := material.Color.Multiply(light.Intensity)
	lightv := light.Position.SubtractPoint(point).Normalize()
	ambient := effectiveColor.MultiplyByScalar(material.Ambient)
	lightDotNormal := lightv.Dot(normalv)
	if lightDotNormal < 0 {
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
