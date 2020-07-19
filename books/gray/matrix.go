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

func (m *Matrix4) multiply(o Matrix4) Matrix4 {
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

func (m *Matrix4) multiplyVector(o Vector) Vector {
	return Vector{
		m[0][0]*o.X + m[0][1]*o.Y + m[0][2]*o.Z + m[0][3],
		m[1][0]*o.X + m[1][1]*o.Y + m[1][2]*o.Z + m[1][3],
		m[2][0]*o.X + m[2][1]*o.Y + m[2][2]*o.Z + m[2][3],
	}
}

func (m *Matrix4) transpose() Matrix4 {
	//todo or implement as loops?
	return Matrix4{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}
}

func (m *Matrix4) determinant() float64 {


	return 0
}

func trimAndParseFloat(s string) float64 {
	s = strings.Trim(s, " ")
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

