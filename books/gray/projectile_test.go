package gray

import (
	"fmt"
	"testing"
)

func Test_tick(t *testing.T) {
	p := projectile{Point{0, 1, 0}, Vector{1, 1, 0}.normalize()}
	e := environment{Vector{0, -0.1, 0}, Vector{-0.01, 0, 0}}
	for p.position.y > 0 {
		p = tick(e, p)
		fmt.Println(p)
	}
}
