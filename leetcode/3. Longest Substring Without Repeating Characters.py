class Solution(object):
    def lengthOfLongestSubstring(self, s):
        result = 0
        current = ''
        for i in s:
            if i in current:
                current = current[current.find(i) + 1:]
            current += i
            result = max(result, len(current))
        return result
