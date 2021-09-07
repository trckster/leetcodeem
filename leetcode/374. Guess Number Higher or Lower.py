# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num: int) -> int:

pick = 1233

def guess(num: int) -> int:
    if num > pick:
        return -1

    if num < pick:
        return 1

    return 0


class Solution:
    def guessNumber(self, n: int) -> int:
        left = 0
        right = n
        pointer = 0
        last_result = 1

        while last_result != 0:
            last_result = guess(pointer)

            if last_result == -1:
                right = pointer
            elif last_result == 1:
                left = pointer + 1

            pointer = left + (right - left) // 2

        return pointer


s = Solution()
print(s.guessNumber(10000))
