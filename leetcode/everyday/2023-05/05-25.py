class Solution:
    def oddString(self, words: List[str]) -> str:
        d = defaultdict(list)
        for s in words:
            t = tuple(ord(b)-ord(a) for a,b in pairwise(s))
            d[t].append(s)
        return next(x[0] for x in d.values() if len(x) == 1)