package MaximumSubarray

func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	cur, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > cur+nums[i] {
			cur = nums[i]
		} else {
			cur = cur + nums[i]
		}

		if cur > max {
			max = cur
		}
	}

	return max
}
