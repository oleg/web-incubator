package figure

import (
	"gray/multid"
	"gray/oned"
)

type Shape interface {
	Transform() multid.Matrix4 //is this method needed in the interface
	Intersect(ray Ray) Inters
	NormalAt(point oned.Point) oned.Vector
	Material() Material
}
