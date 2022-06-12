class Solution:
    def maximumUniqueSubarray(self, nums):
        used = set()

        answer = sum = left = right = 0

        while right < len(nums):
            sum += nums[right]

            if nums[right] in used:
                while nums[left] != nums[right]:
                    used.remove(nums[left])
                    sum -= nums[left]
                    left += 1
                sum -= nums[left]
                left += 1
            else:
                used.add(nums[right])

            answer = max(answer, sum)
            right += 1

        return answer


s = Solution()
print(s.maximumUniqueSubarray([4, 2, 4, 5, 6]))
print(s.maximumUniqueSubarray([5, 2, 1, 2, 5]))
