
class Solution:
    # 跳跃游戏, 视频拼接
    def minTaps(self, n: int, ranges: List[int]) -> int:
        # 求当前节点能跳跃的最远距离
        right_most = [0] * (n+1)
        for i, r in enumerate(ranges):
            left = max(i-r, 0)
            right_most[left] = max(right_most[left], i + r)
        
        ans = 0
        cur = 0
        next = 0
        for i in range(n):  # n已经是终点了，不需要考虑
            next = max(next, right_most[i])
            if i == cur:
                if i == next:
                    return -1
                cur = next
                ans += 1
        return ans
            
        