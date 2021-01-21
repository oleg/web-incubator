package corm

//2.3.1
func MergeSort(arr []int) []int {
	l := len(arr)
	if l == 0 || l == 1 {
		return arr
	}

	h := l / 2
	a := MergeSort(arr[:h])
	b := MergeSort(arr[h:])
	return merge(a, b)
}

func merge(a []int, b []int) []int {
	res := make([]int, len(a)+len(b))
	ai := 0
	bi := 0
	for ai < len(a) && bi < len(b) {
		if a[ai] <= b[bi] {
			res[ai+bi] = a[ai]
			ai++
		} else {
			res[ai+bi] = b[bi]
			bi++
		}
	}
	for ; ai < len(a); ai++ {
		res[ai+bi] = a[ai]
	}
	for ; bi < len(b); bi++ {
		res[ai+bi] = b[bi]
	}
	return res
}
