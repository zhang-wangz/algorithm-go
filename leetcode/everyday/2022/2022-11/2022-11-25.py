class Solution:
    def expressiveWords(self, s: str, words: List[str]) -> int:
        def helper(s: str, w: str) -> bool:
            n, wn  = len(s), len(w)
            if wn > n: return False
            i = j = 0
            while i < n and j < wn:
                if s[i] != w[j]:
                    return False
                ch = s[i]
                cnti = 0
                while i < n and s[i] == ch:
                    cnti += 1
                    i += 1
                cntj = 0
                while j < wn and w[j] == ch:
                    cntj += 1
                    j += 1
                if cnti < cntj:
                    return False
                if cnti > cntj and cnti < 3:
                    return False
            return i == n and j == wn
        return sum([int(helper(s, w)) for w in words])