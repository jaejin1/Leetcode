class Solution:
    def isValid(self, s: str) -> bool:
        dic = {')': '(', ']': '[', '}': '{'}
        stack = []
        if len(s) == 1:
            return False
        for i in s:
            if i == '(' or i == '[' or i == '{':
                stack.append(i)
            else:
                if len(stack) > 0 and dic[i] == stack[len(stack)-1]:
                    stack = stack[:-1]
                else:
                    return False
        
        if len(stack) == 0:
            return True
        return False
    