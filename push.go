package collection

func Push(a []int, v int) ([]int, int) {
	return append(a, v), len(a)
}
