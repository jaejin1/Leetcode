package ProductofArrayExceptSelf

func productExceptSelf(nums []int) []int {
	resultA := make([]int, len(nums))
	resultB := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			resultA[i] = 1
		} else {
			resultA[i] = resultA[i-1] * nums[i-1]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			resultB[i] = 1
		} else {
			resultB[i] = resultB[i+1] * nums[i+1]
		}
	}

	for i, n := range resultA {
		resultB[i] *= n
	}
	return resultB
}
