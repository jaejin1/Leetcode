class Solution:
    def countSubstrings(self, s: str) -> int:
        dp = [[False for i in range(len(s))] for j in range(len(s))]
        
        result = 0
        for i in range(0, len(s)):
            for j in range(i, -1, -1):
                
                if s[i] == s[j] and (i-j <= 2 or dp[i-1][j+1]):
                    
                    dp[i][j] = True
                    result += 1

        return result
