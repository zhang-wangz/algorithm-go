# -*- coding: utf-8 -*-
# @Time    : 2025/1/2 14:20 
# @File    : 01-01.py

class Solution:
    def convertDateToBinary(self, date: str) -> str:
        data = date.split('-')
        return '-'.join(bin(int(c))[2:] for c in data)

