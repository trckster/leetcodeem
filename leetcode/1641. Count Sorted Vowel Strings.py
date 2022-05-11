class Solution:
    # https://en.wikipedia.org/wiki/Pentatope_number
    def countVowelStrings(self, n):
        return (n + 1) * (n + 2) * (n + 3) * (n + 4) // 24


s = Solution()
print(s.countVowelStrings(1))
print(s.countVowelStrings(2))
print(s.countVowelStrings(3))
print(s.countVowelStrings(4))
print(s.countVowelStrings(5))
print(s.countVowelStrings(6))
print(s.countVowelStrings(7))
print(s.countVowelStrings(8))
print(s.countVowelStrings(9))
print(s.countVowelStrings(10))
print(s.countVowelStrings(33))
print(s.countVowelStrings(50))
