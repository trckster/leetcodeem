class Solution:
    def romanToInt(self, s: str) -> int:
        result = 0

        patterns = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }
        inverted_patterns = {
            'IV': 4,
            'IX': 9,
            'XL': 40,
            'XC': 90,
            'CD': 400,
            'CM': 900
        }

        for key in inverted_patterns:
            if key in s:
                result += inverted_patterns[key]
                s = s.replace(key, '')

        for c in s:
            result += patterns[c]

        return result


# solution = Solution()
# print(solution.romanToInt('LVIII'))
