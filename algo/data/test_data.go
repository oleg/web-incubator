package data

type Test struct {
	Name     string
	Array    []int
	Expected []int
}

func TestData() []Test {
	return []Test{
		{"empty", []int{}, []int{}},
		{"one element", []int{7}, []int{7}},
		{"two elements", []int{7, 3}, []int{3, 7}},
		{"ten elements", []int{2, 3, 1, 4, 1, 46, 7, -8, 0, 5}, []int{-8, 0, 1, 1, 2, 3, 4, 5, 7, 46}},
	}
}

func ReverseTestData() []Test {
	return []Test{
		{"empty", []int{}, []int{}},
		{"one element", []int{7}, []int{7}},
		{"two elements", []int{7, 3}, []int{7, 3}},
		{"ten elements", []int{2, 3, 1, 4, 1, 46, 7, -8, 0, 5}, []int{46, 7, 5, 4, 3, 2, 1, 1, 0, -8}},
	}
}
