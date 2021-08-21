package SingleNumber

func singleNumber(nums []int) int {
	result := 0
	for _, s := range nums {
		result ^= s
	}

	return result
}
