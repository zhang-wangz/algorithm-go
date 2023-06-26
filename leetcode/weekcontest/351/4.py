

class Solution:
    def survivedRobotsHealths(self, positions: List[int], healths: List[int], directions: str) -> List[int]:
        st = []
        healths = ((c, i) for i, c in enumerate(healths))
        for i, (p, h, d) in enumerate(sorted(zip(positions, healths, directions), key=lambda x: x[0])):
            cur = [h[1], h[0], d]
            # 健康序号， 当前健康度， 方向
            while st and st[-1][2] == 'R' and cur[2] == 'L' and cur[1] > 0:
                if st[-1][1] < cur[1]:
                    cur[1] = cur[1] - 1
                    st.pop()
                elif st[-1][1] > cur[1]:
                    last = st.pop()
                    st.append((last[0], last[1]-1, last[2]))
                    cur[1] = 0
                else:
                    st.pop()
                    cur[1] = 0
            if cur[1] > 0:
                st.append((cur[0], cur[1], cur[2]))
        ans = [0] * len(positions)
        while st:
            l = st.pop()
            ans[l[0]] = l[1]
        return [i for i in ans if i > 0]

                
        