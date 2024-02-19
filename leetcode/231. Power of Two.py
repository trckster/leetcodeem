import math


class Solution(object):
    def isPowerOfTwo(self, n):
        if n - math.floor(n) != 0 or n == 0:
            return False

        if math.floor(n) == 1:
            return True

        return self.isPowerOfTwo(n / 2)
