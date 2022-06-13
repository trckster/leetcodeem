class Solution:
    def minimumTotal(self, triangle):
        for i in range(1, len(triangle)):
            for j in range(len(triangle[i])):
                if j == 0:
                    triangle[i][j] += triangle[i - 1][j]
                elif j == len(triangle[i]) - 1:
                    triangle[i][j] += triangle[i - 1][j - 1]
                else:
                    left = triangle[i - 1][j - 1]
                    right = triangle[i - 1][j]
                    triangle[i][j] += min(left, right)

        return min(triangle[-1])


s = Solution()

data1 = [[2],[3,4],[6,5,7],[4,1,8,3]]

print(s.minimumTotal(data1))
