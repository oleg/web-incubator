package corm

//2.2-2
func SelectionSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		mini := i
		min := arr[mini]
		for j := i; j < l; j++ {
			if arr[j] < min {
				min = arr[j]
				mini = j
			}
		}
		arr[i], arr[mini] = arr[mini], arr[i]
	}
}
