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
//	return p.position.addVector(p.velocity)
//}

func (proj projectile) tick(env environment) projectile {
	position := proj.position.addVector(proj.velocity)
	velocity := proj.velocity.addVector(env.gravity).addVector(env.wind)
	return projectile{position, velocity}
}
