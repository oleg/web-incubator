package stepik

//func main() {
//	var n int
//	fmt.Scan(&n)
//	ks := kadd(n)
//	fmt.Println(len(ks))
//	fmt.Print(ks)
//}

func kadd(n int) []int {
	answer := make([]int, 0)
	for res, v := n, 1; res != 0; v++ {
		dec := res - v
		if dec == 0 || dec > v {
			res = dec
			answer = append(answer, v)
		}
	}
	return answer
}
