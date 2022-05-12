class Solution:
    def __init__(self):
        self.answer = []

    def count_ways(self, prefix, remaining):
        if len(remaining) == 0:
            self.answer.append(prefix)
            return

        next_applicants = set(remaining)

        for i in next_applicants:
            i_index = remaining.index(i)
            next_remaining = remaining[:i_index] + remaining[i_index + 1:]
            p = prefix.copy()
            p.append(i)
            self.count_ways(p, next_remaining)

    def permuteUnique(self, nums):
        self.answer = []

        self.count_ways([], nums)

        return self.answer


s = Solution()
print(s.permuteUnique([1, 1, 2]))
print(s.permuteUnique([1, 2, 3]))
print(s.permuteUnique([1, 2, 3, 4, 5, 6, 7, 8]))
