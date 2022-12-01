class Solution:
    def maxRotateFunction(self, nums: list) -> int:
        together = sum(nums)

        f_x_previous = answer = sum([i * nums[i] for i in range(len(nums))])

        i = len(nums) - 1
        while i >= 0:
            f_x = f_x_previous + together - nums[i] * len(nums)

            answer = max(answer, f_x)

            f_x_previous = f_x

            i -= 1

        return answer
