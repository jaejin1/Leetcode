class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        l = self.loop(nums,target, True)
        
        if l == len(nums) or nums[l] != target:
            return [-1, -1]
    
        r = self.loop(nums, target, False)
        
        return [l, r - 1]
        
    def loop(self, nums: List[int], target: int, left: bool) -> int:
        l = 0
        r = len(nums)
        
        while l < r:
            mid = (l + r) // 2
            
            if nums[mid] > target or (left and target == nums[mid]):
                r = mid
            else:
                l = mid + 1
                
        return l