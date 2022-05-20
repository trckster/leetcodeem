class Solution:
    def __init__(self):
        self.cache = []
        self.grid = []
        self.height = 0
        self.width = 0

    def prepare_cache(self, grid):
        self.cache = []
        for i in range(len(grid)):
            self.cache.append([-1 for i in range(len(grid[i]))])
        self.grid = grid
        self.height = len(grid)
        self.width = len(grid[0])

    def get_count_of_ways(self, x, y):
        if x < 0 or y < 0:
            return 0

        if self.cache[x][y] != -1:
            return self.cache[x][y]

        if self.grid[x][y] == 1:
            return 0

        if x == 0 and y == 0:
            return 1

        ways_to_come_from_left = self.get_count_of_ways(x, y - 1)
        ways_to_come_from_top = self.get_count_of_ways(x - 1, y)

        result = ways_to_come_from_top + ways_to_come_from_left

        self.cache[x][y] = result

        return result

    def uniquePathsWithObstacles(self, grid):
        self.prepare_cache(grid)

        if grid[self.height - 1][self.width - 1] == 1:
            return 0

        return self.get_count_of_ways(self.height - 1, self.width - 1)


s = Solution()
print(s.uniquePathsWithObstacles([[0, 1], [0, 0]]))
print(s.uniquePathsWithObstacles([[1, 0], [0, 0]]))
print(s.uniquePathsWithObstacles([[0, 0], [1, 0]]))
print(s.uniquePathsWithObstacles([[0, 0], [0, 1]]))
