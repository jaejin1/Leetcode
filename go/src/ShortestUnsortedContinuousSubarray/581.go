package ShortestUnsortedContinuousSubarray

import "sort"

func findUnsortedSubarray(nums []int) int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(nums)

	var count int
	for i, _ := range temp {
		if temp[i] != nums[i] {
			break
		} else {
			count++
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] != nums[i] {
			break
		} else {
			count++
		}
	}

	if len(nums) == count/2 {
		return 0
	}
	return len(nums) - count
}
