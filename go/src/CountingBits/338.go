package CountingBits

func countBits(num int) []int {
	result := make([]int, num+1)

	if num == 0 {
		return result
	}

	for i := 1; i < num+1; i++ {
		if i%2 == 1 {
			result[i] = result[i/2] + 1
		} else {
			result[i] = result[i/2]
		}
	}
	return result
}
