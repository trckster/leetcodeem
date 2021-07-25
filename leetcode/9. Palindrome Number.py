class Solution:
    def isPalindrome(self, x: int) -> bool:
        s = str(x)

        i = 0
        while i < len(s):
            if s[i] != s[len(s) - 1 - i]:
                return False
            i += 1
        return True


solution = Solution()
print(solution.isPalindrome(1233231))
