class Solution:
    def __init__(self):
        self.mapping = {
            '2': 'abc',
            '3': 'def',
            '4': 'ghi',
            '5': 'jkl',
            '6': 'mno',
            '7': 'pqrs',
            '8': 'tuv',
            '9': 'wxyz'
        }

    def assemble_strings(self, prefixes, to_add):
        if len(to_add) == 0:
            return prefixes

        current_chars = self.mapping[to_add[0]]

        result = []

        for prefix in prefixes:
            new_prefixes = []

            for c in current_chars:
                new_prefixes.append(prefix + c)

            result += self.assemble_strings(new_prefixes, to_add[1:])

        return result

    def letterCombinations(self, digits):
        if len(digits) == 0:
            return []

        return self.assemble_strings([''], digits)


s = Solution()
print(s.letterCombinations('234'))
