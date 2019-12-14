class Solution(object):
    def letterCasePermutation(self, S):
        result = []
        self.dfs(S, 0, "", result)
        return result
    
    def dfs(self, S, index, item, result):
        if len(item) == len(S):
            result.append(item)
            return
        
        for i in range(index, len(S)):
            if S[i] not in "0123456789":
                self.dfs(S, i+1, item + S[i].upper(), result)
                self.dfs(S, i+1, item + S[i].lower(), result)
            else:
                self.dfs(S, i+1, item + S[i], result)
