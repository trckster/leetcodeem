class Solution:
    def sortArrayByParity(self, nums):
        nums.sort(key=lambda x: (x % 2))

        return nums


s = Solution()
print(s.sortArrayByParity([6, 3, 4, 1, 2, 3, 3, 10, 21]))

