class Solution:
    def jump(self, nums: list) -> int:
        jumps = [-1 for i in nums]
        jumps[0] = 0

        for i in range(len(nums) - 1):
            jump_length = nums[i]
            furthest_reachable = i + jump_length

            if furthest_reachable + 1 >= len(nums):
                return jumps[i] + 1

            while jumps[furthest_reachable] == -1:
                jumps[furthest_reachable] = jumps[i] + 1
                furthest_reachable -= 1

        return jumps[-1]