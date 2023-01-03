from typing import List

class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        res = 0
        for num in sorted(boxTypes, key=lambda x:-x[1]):
            cnt, n = num
            if truckSize >= cnt:
                res += cnt * n
                truckSize -= cnt
            elif truckSize > 0 and truckSize < cnt:
                res += truckSize * n
                truckSize = 0
            if  truckSize <= 0 : break
        return res
            