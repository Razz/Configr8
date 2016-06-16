package plugin

func Times(n int) []int {
	n += 1
	sliceInt := make([]int, n)
	for i := 0; i <= len(sliceInt); i++ {
		sliceInt[i] = i
	}
	return sliceInt
}
