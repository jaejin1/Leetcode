class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_position = {}
        for i, n in enumerate(nums):
            if target - n in nums_position:
                return [nums_position[target-n], i]
            nums_position[n] = i

        return None