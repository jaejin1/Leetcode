package FindFirstandLastPositionofElementinSortedArray

func searchRange(nums []int, target int) []int {

	l := loop(nums, target, true)

	if l == len(nums) || nums[l] != target {
		return []int{-1, -1}
	}

	r := loop(nums, target, false)

	return []int{l, r - 1}
}

func loop(nums []int, target int, left bool) int {
	l := 0
	r := len(nums)

	for l < r {
		mid := (l + r) / 2

		if nums[mid] > target || (left && target == nums[mid]) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
