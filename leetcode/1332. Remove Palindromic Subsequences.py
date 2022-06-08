class Solution:
    def removePalindromeSub(self, s):
        def is_palindrome():
            i = 0

            while i < len(s) - i - 1:
                if s[i] != s[len(s) - i - 1]:
                    return False
                i += 1

            return True

        if is_palindrome():
            return 1
        else:
            return 2
