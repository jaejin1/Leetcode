package SingleNumber

func singleNumber(nums []int) int {
	set := make(map[int]int, len(nums))
	result := 0

	for _, s := range nums {
		_, ok := set[s]
		if ok == true {
			delete(set, s)
			result -= s
		} else {
			set[s] = 1
			result += s
		}
	}

	return result
}
