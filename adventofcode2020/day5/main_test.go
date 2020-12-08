package main

import (
	"testing"
)

func Test_day5_task1_row_column(t *testing.T) {
	tests := []struct {
		code           string
		expectedRow    int
		expectedColumn int
	}{
		{"BFFFBBFRRR", 70, 7},
		{"FFFBBBFRRR", 14, 7},
		{"BBFFBBFRLL", 102, 4},
	}
	for _, test := range tests {
		t.Run(test.code, func(t *testing.T) {

			row, column := parseRowAndColumn(test.code)
			if column != test.expectedColumn || row != test.expectedRow {
				t.Errorf("Wrong row %d != %d or column %d != %d or",
					row, test.expectedRow, column, test.expectedColumn)
			}
		})
	}
}

func Test_day5_seat_id(t *testing.T) {
	tests := []struct {
		name     string
		row      int
		column   int
		expected int
	}{
		{"seat 567", 70, 7, 567},
		{"seat 119", 14, 7, 119},
		{"seat 820", 102, 4, 820},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			seatId := seatId(test.row, test.column)
			if seatId != test.expected {
				t.Errorf("Wrong seatId %d", seatId)
			}
		})
	}
}


