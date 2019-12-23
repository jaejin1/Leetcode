class Solution(object):
    def climbStairs(self, n):
        if n == 1:
            return 1
        tmp, result = 1, 2
        for i in range(2, n):
            tmp2 = result
            result = tmp+result
            tmp = tmp2
        return result
