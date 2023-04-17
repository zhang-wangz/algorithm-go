DAYS_SUM = list(accumulate((31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31), initial=0))

def calc_days(date: str) -> int:
    return DAYS_SUM[int(date[:2]) - 1] + int(date[3:])

class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        left = calc_days(max(arriveAlice, arriveBob))  # 直接比较字符串再计算
        right = calc_days(min(leaveAlice, leaveBob))
        return max(right - left + 1, 0)  # 答案不能为负数
