package plugin

func Times(n int) []int {
	var sliceInt []int
	for i := 1; i != n; i++ {
		sliceInt = append(sliceInt, i)
	}
	return sliceInt
}
