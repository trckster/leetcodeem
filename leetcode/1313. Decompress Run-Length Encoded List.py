# from typing import List


class Solution:
    def decompressRLElist(self, nums: List[int]) -> List[int]:
        result = []

        i = 1
        while i < len(nums):
            result += [nums[i] for j in range(nums[i - 1])]
            i += 2

        return result


# solution = Solution()
# print(solution.decompressRLElist([1, 1, 2, 3]))
