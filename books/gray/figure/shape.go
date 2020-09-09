package figure

import "gray/multid"

type Shape interface {
	Transform() multid.Matrix4 //is this method needed in the interface
	Intersect(ray Ray) Inters
}
