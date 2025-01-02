class Solution:
    # hashmap
    def getFolderNames(self, names: List[str]) -> List[str]:
        index = {}
        ans = []
        for name in names:
            if name not in index:
                ans.append(name)
                index[name] = 1
            else:
                idx = index[name]
                while name+'('+str(idx)+')' in index:
                    idx+=1
                t = name+'('+str(idx)+')'
                ans.append(t)
                index[name] = idx+1
                index[t] = 1
        return ans