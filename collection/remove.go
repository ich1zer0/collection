package collection

func Remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...)
}
