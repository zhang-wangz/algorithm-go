class Solution:
    def checkArithmeticSubarrays(self, nums: List[int], l: List[int], r: List[int]) -> List[bool]:
        ans = []
        for l, r  in zip(l, r):
            maxv = max(nums[l:r+1])
            minv = min(nums[l:r+1])

            if maxv == minv:
                ans.append(True)
                continue
            if (maxv - minv) % (r-l) != 0:
                ans.append(False)
                continue
            
            d = (maxv - minv) // (r-l)
            vis = set()
            flag = True
            for i in range(l, r+1):
                if (nums[i]-minv) % d != 0:
                    flag = False
                    break
                t = (nums[i]-minv) // d
                if t in vis:
                    flag = False
                    break
                vis.add(t)
            ans.append(flag)
        return ans