class Solution:
    def hammingWeight(self, n: int) -> int:
        answer = 0
        while n:
            if n & 1:
                answer += 1
            n >>= 1
        return answer
