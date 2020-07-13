package gray

import (
	"log"
	"strconv"
	"strings"
)

type Matrix4 [4][4]float64

func (m Matrix4) multiply(o Matrix4) Matrix4 {
	return Matrix4{}
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

//func NewMatrix4(
//	e00, e01, e02, e03,
//	e10, e11, e12, e13,
//	e20, e21, e22, e23,
//	e30, e31, e32, e33 float64) Matrix4 {
//	return Matrix4{
//		[4]float64{e00, e01, e02, e03},
//		[4]float64{e10, e11, e12, e13},
//		[4]float64{e20, e21, e22, e23},
//		[4]float64{e30, e31, e32, e33},
//	}
//}
