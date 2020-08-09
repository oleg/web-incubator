package sort

type Test struct {
	name     string
	arr      []int
	expected []int
}

var TestData = []Test{
	{"empty", []int{}, []int{}},
	{"one element", []int{7}, []int{7}},
	{"two elements", []int{7, 3}, []int{3, 7}},
	{"ten elements", []int{2, 3, 1, 4, 1, 46, 7, -8, 0, 5}, []int{-8, 0, 1, 1, 2, 3, 4, 5, 7, 46}},
}
var ReverseTestData = []Test{
	{"empty", []int{}, []int{}},
	{"one element", []int{7}, []int{7}},
	{"two elements", []int{7, 3}, []int{7, 3}},
	{"ten elements", []int{2, 3, 1, 4, 1, 46, 7, -8, 0, 5}, []int{46, 7, 5, 4, 3, 2, 1, 1, 0, -8}},
}
