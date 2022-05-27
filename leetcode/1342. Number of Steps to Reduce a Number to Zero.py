class Solution:
    def numberOfSteps(self, num):
        if not num:
            return 0

        return (self.numberOfSteps(num - 1) if num % 2 else self.numberOfSteps(num // 2)) + 1
