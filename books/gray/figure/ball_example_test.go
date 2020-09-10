package figure

import (
	"gray/multid"
	"gray/oned"
	"testing"
)

//todo do not use example in the name
func Test_ball_example(t *testing.T) {
	rayOrigin := oned.Point{0, 0, - 5}
	wallSize := 7.
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2.
	canvas := multid.NewCanvas(canvasPixels, canvasPixels)
	red := oned.Color{1, 0, 0}
	transform := multid.Shearing(1, 0, 0, 0, 0, 0).Multiply(multid.Scaling(0.5, 1, 1))
	sphere := MakeSphereT(transform)

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := oned.Point{worldX, worldY, 10}
			ray := Ray{rayOrigin, position.SubtractPoint(rayOrigin).Normalize()}
			if hit, _ := sphere.Intersect(ray).Hit(); hit {
				canvas.Pixels[x][y] = red
			}
		}
	}

	canvas.MustToPNG("ball_example_test.png")
}
