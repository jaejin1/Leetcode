class Solution:
    def rob(self, nums: List[int]) -> int:
        now, last = 0, 0
        # last = f(k-2)
        # now = f(k-1)
        for i in nums:
            last, now = now, max(last + i, now)
                
        return(now)
                