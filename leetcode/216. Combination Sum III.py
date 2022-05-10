class Solution:
    def __init__(self):
        self.answer = []
        self.goal = None

    def add_combinations(self, start, remaining):
        try:
            last_number = start[-1]
        except IndexError:
            last_number = 0

        minimum_remaining_sum = remaining * (remaining + 1) // 2 + (remaining * last_number)
        maximum_remaining_sum = minimum_remaining_sum + (remaining * (9 - remaining - last_number))
        existing_sum = sum(start)

        if self.goal < minimum_remaining_sum + existing_sum:
            return
        elif self.goal > maximum_remaining_sum + existing_sum:
            return

        if remaining == 0:
            self.answer.append(start)
            return

        for i in range(last_number + 1, 10 - remaining + 1):
            new_array = start.copy()

            new_array.append(i)

            self.add_combinations(new_array, remaining - 1)

    def combinationSum3(self, k: int, n: int):
        self.goal = n

        self.add_combinations([], k)

        return self.answer


s = Solution()
print(s.combinationSum3(2, 5))
