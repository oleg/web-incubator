package stepik

func fib(n int) int {
	if n <= 1 {
		return n
	}
	n0 := 0
	n1 := 1
	var n2 int
	for i := 0; i < n-1; i++ {
		n2 = n1 + n0
		n1, n0 = n2, n1

	}
	return n2
}

func fibLd(n int) int {
	if n <= 1 {
		return n
	}
	n0 := 0
	n1 := 1
	var n2 int
	for i := 0; i < n-1; i++ {
		n2 = (n1 + n0) % 10
		n1, n0 = n2, n1

	}
	return n2
}

func fibBigMod(n, m int) int {
	if n <= 1 {
		return n
	}
	n0 := 0
	n1 := 1
	var n2 int
	rem := make([]int, 0, 100)
	rem = append(rem, 0, 1)
	for i := 0; i <= 6*m; i++ {
		n2 = (n1 + n0) % m
		n1, n0 = n2, n1
		rem = append(rem, n2)
		if n0 == 0 && n1 == 1 {
			break
		}
	}
	return rem[n%(len(rem)-2)]
}
