class Solution:
    def distributeCandies(self, candies):
        candies.sort()

        answer = 1
        for i in range(1, len(candies)):
            if candies[i] != candies[i - 1]:
                answer += 1

        return min(answer, len(candies) // 2)


s = Solution()
print(s.distributeCandies([3, 3, 1, 1, 2, 2, 3, 3]))
