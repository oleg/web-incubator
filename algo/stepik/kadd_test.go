package stepik

import "testing"

func Test_kadd(t *testing.T) {
	tests := []struct {
		name string
		n    int
		ks   []int
	}{
		{"case 1", 1, []int{1}},
		{"case 2", 2, []int{2}},
		{"case 3", 3, []int{1, 2}},
		{"case 4", 4, []int{1, 3}},
		{"case 6", 6, []int{1, 2, 3}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			ks := kadd(test.n)

			if len(ks) != len(test.ks) {
				t.Errorf("Wrong len(ks) %v, expected %v", ks, test.ks)
			}
			for i, v := range test.ks {
				if v != test.ks[i] {
					t.Errorf("Wrong %d elemnt %d expected %d", i, v, test.ks[i])
				}
			}
		})
	}
}
