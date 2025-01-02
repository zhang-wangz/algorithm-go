# -*- coding: utf-8 -*-
# @Time    : 2025/1/2 14:17 
# @File    : 01-02.py
# https://leetcode.cn/problems/my-calendar-i/?envType=daily-question&envId=2025-01-02
from bisect import *
from math import inf


class MyCalendar:

    def __init__(self):
        self.cal = [[-1, -1], [inf, inf]]

    def book(self, startTime: int, endTime: int) -> bool:
        i = bisect_right(self.cal, [startTime, endTime])
        if startTime < self.cal[i - 1][1] or endTime > self.cal[i][0]: return False
        self.cal.insert(i, [startTime, endTime])
        return True

# Your MyCalendar object will be instantiated and called as such:
# obj = MyCalendar()
# param_1 = obj.book(startTime,endTime)
