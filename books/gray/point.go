package gray

type Point Tuple //w = 1

func (t Point) addVector(o Vector) Point {
	return Point(Tuple(t).add(Tuple(o)))
}

func (t Point) subtractVector(o Vector) Point {
	return Point(Tuple(t).subtract(Tuple(o)))
}

func (t Point) subtractPoint(o Point) Vector {
	return Vector(Tuple(t).subtract(Tuple(o)))
}
