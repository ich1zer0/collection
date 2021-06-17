package collection

func Insert(a []int, v int, p int) []int {
	a = append(a, 0)
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}
