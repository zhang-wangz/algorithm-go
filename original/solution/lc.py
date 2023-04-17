
from heapq import heappop, heappush

def solve(cakes, limits):
    ans = 0
    hq = []
    day = 0
    heappush(hq, (day+limits[day], cakes[day], day))
    while hq:
        while hq:
            q = heappop(hq)
            if q[0] > day:
                print(day, q)
                ans += 1
                cnt = q[1] - 1
                if cnt > 0: heappush(hq, (q[0], cnt, q[2]))
                break
        day += 1
        if day < len(cakes): heappush(hq, (day+limits[day], cakes[day], day))
    return ans

print(solve([1,3,2],[3,5,2]))
        