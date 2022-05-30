import random


class Solution:
    def count_fills(self, n_str, divisor: int):
        n = int(n_str)
        result = 0
        sum = 0

        while divisor + sum <= n:
            sum += divisor
            result += 1

        return result, sum

    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == 0:
            return 0

        answer = '0'
        if dividend > 0 > divisor or dividend < 0 < divisor:
            answer = '-0'

        dividend = str(abs(dividend))
        divisor = str(abs(divisor))

        indent = 0
        while indent + len(divisor) <= len(dividend):
            left = dividend[:len(divisor) + indent]
            left_length = len(left)
            right = dividend[len(divisor) + indent:]
            count, to_subtract = self.count_fills(left, int(divisor))
            answer += str(count)
            new_left = int(left) - to_subtract
            new_left = str(new_left).rjust(left_length, '0')
            dividend = new_left + right
            indent += 1

        answer = min(int(answer), 2147483647)
        return max(answer, -2147483648)


s = Solution()

for i in range(1000):
    a = random.randint(-1000, 1000)
    b = random.randint(-1000, 1000)
    expected = int(a / b)
    given = s.divide(a, b)

    if expected == given:
        print('.', end='')
    else:
        print(f'Problem: {a} / {b}? Expected: {expected}, given: {given}')
        exit(0)
