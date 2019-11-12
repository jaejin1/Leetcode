package Subsets

func subsets(nums []int) [][]int {
	var result [][]int

	for i := 0; i < int(1<<uint(len(nums))); i++ {
		var temp []int
		for j, n := range nums {
			if i&(1<<uint(j)) != 0 {
				temp = append(temp, n)
			}
		}

		result = append(result, temp)
	}
	return result
}
