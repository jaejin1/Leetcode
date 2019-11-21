class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        if not nums:
            return 0
        
        dic = {}
        for i in nums:
            if i not in dic:
                dic[i] = 1
            else:
                dic[i] = dic[i] + 1
        
        for k, v in dic.items():
            if len(nums) / 2 <= v:
                return k

        return 0