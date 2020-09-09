package oned

type projectile struct {
	position Point
	velocity Vector
}
type environment struct {
	gravity Vector
	wind    Vector
}

//func (p projectile) move() (position Point) {
//	return p.position.AddVector(p.velocity)
//}

func (proj projectile) tick(env environment) projectile {
	position := proj.position.AddVector(proj.velocity)
	velocity := proj.velocity.addVector(env.gravity).addVector(env.wind)
	return projectile{position, velocity}
}
