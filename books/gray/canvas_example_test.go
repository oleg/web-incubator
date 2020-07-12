package gray

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	start := Point{0, 1, 0}
	velocity := Vector{1, 1.8, 0}.normalize().multiplyScalar(11.25)
	p := projectile{start, velocity}

	gravity := Vector{0, -0.1, 0}
	wind := Vector{-0.01, 0, 0}
	e := environment{gravity, wind}

	c := NewCanvas(900, 500)

	for p.position.X >= 0 && p.position.Y > 0 {
		x := int(p.position.X)
		y := int(p.position.Y)
		c.Pixels[x][y] = NewColor(1, 0, 0)
		p = p.tick(e)
	}

	err := c.toPNG("canvas_example_test.png")
	if err != nil {
		log.Panic(err)
	}

}
