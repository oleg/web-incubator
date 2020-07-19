package gray

import (
	"log"
	"strconv"
	"strings"
)

type Matrix4 [4][4]float64

var IdentityMatrix Matrix4 = [4][4]float64{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

func (m Matrix4) multiply(o Matrix4) Matrix4 {
	r := Matrix4{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				r[i][j] += m[i][k] * o[k][j]
			}
		}
	}
	return r
}

func (m Matrix4) multiplyVector(o Vector) Vector {
	return Vector{
		m[0][0]*o.X + m[0][1]*o.Y + m[0][2]*o.Z + m[0][3],
		m[1][0]*o.X + m[1][1]*o.Y + m[1][2]*o.Z + m[1][3],
		m[2][0]*o.X + m[2][1]*o.Y + m[2][2]*o.Z + m[2][3],
	}
}

//todo must?
func NewMatrix4(str string) Matrix4 {
	m := Matrix4{}
	rows := strings.Split(str, "\n")
	if len(rows) != 4 {
		log.Fatal("must have 4 rows")
	}
	for i, row := range rows {
		columns := strings.Split(row, "|")
		if len(columns) != 6 {
			log.Fatal("must have 6 separators for 4 columns")
		}
		for j, s := range columns[1:5] {
			m[i][j] = trimAndParseFloat(s)
		}
	}
	return m
}

func trimAndParseFloat(s string) float64 {
	s = strings.Trim(s, " ")
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

//how to reuse this?
func multiply(a, b [][]float64) [][]float64 {
	if len(a) == 0 && len(b) == 0 {
		return nil
	}

	aw := len(a)
	ah := len(a[0])

	bw := len(b)
	bh := len(b[0])

	if ah != bw {
		panic("not possible 2") //todo test
	}

	c := make([][]float64, aw)
	for i := 0; i < aw; i++ {
		c[i] = make([]float64, bh)
	}

	for i := 0; i < aw; i++ {
		for j := 0; j < bh; j++ {
			for k := 0; k < ah /*&& k < bw*/ ; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
