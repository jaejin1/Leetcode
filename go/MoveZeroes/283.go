package MoveZeroes

func moveZeroes(nums []int) {
	zero := 0
	for i, n := range nums {
		if n != 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
		}
	}
}
