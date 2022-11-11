class Solution:
    def halvesAreAlike(self, s: str) -> bool:
        lst = ['a','e','i','o','u','A','E','I','O','U']
        mid = len(s) // 2
        return sum(x in lst for x in s[0:mid]) == sum(x in lst for x in s[mid:])