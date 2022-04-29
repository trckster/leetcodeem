class Solution:
    def getPriceToReach(self, index):
        if index in self.cache:
            return self.cache[index]

        result = min(
            self.getPriceToReach(index - 2) + self.costs[index - 2],
            self.getPriceToReach(index - 1) + self.costs[index - 1],
        )

        self.cache[index] = result

        return result

    def minCostClimbingStairs(self, costs):
        self.costs = costs + [0]
        self.cache = {
            0: 0,
            1: 0
        }

        return self.getPriceToReach(len(costs))


s = Solution()
print(s.minCostClimbingStairs([1, 100, 1, 1, 1, 100, 1, 1, 100, 1]))
