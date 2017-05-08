package plugin

// add
func Add(n ...int) int {
	var sum int
	for i := range n {
		sum += i
	}
	return sum
}
