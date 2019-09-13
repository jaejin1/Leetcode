class Solution:
    def findAndReplacePattern(self, words: List[str], pattern: str) -> List[str]:
        def match(word):
            m1, m2 = {}, {}
            for w, p in zip(word, pattern):
                print(w, p)
                if w not in m1:
                    m1[w] = p
                if p not in m2:
                    m2[p] = w
                    
                #print(m1, m2)
                if (m1[w], m2[p]) != (p, w):
                    return False
            
            return True

        return filter(match, words)