package collection

func Pop(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}
