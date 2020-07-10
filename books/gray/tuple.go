package gray

type Tuple struct {
	x, y, z float64
}

func (t Tuple) add(o Tuple) Tuple {
	return Tuple{t.x + o.x, t.y + o.y, t.z + o.z}
}

func (t Tuple) subtract(o Tuple) Tuple {
	return Tuple{t.x - o.x, t.y - o.y, t.z - o.z}
}

func (t Tuple) negate() Tuple {
	return Tuple{- t.x, - t.y, - t.z}
}

func (t Tuple) multiply(scalar float64) Tuple {
	return Tuple{t.x * scalar, t.y * scalar, t.z * scalar}
}

func (t Tuple) divide(scalar float64) Tuple {
	return Tuple{t.x / scalar, t.y / scalar, t.z / scalar}
}

/*
​ 	​constant​ EPSILON ← 0.00001
​
​ 	​function​ equal(a, b)
​ 	  ​if​ abs(a - b) < EPSILON
​ 	    ​return​ true
​ 	  ​else​
​ 	    ​return​ false
​ 	  ​end​ ​if​
​ 	​end​ ​function
*/
