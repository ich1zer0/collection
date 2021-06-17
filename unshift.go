package collection

func Unshift(a []int, v int) []int {
	return append([]int{v}, a...)
}
