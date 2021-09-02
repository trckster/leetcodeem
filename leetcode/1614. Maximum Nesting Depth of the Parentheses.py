class Solution:
    def maxDepth(self, s: str) -> int:
        depth = 0
        answer = 0

        for c in s:
            if c == '(':
                depth += 1
                answer = max(depth, answer)
            elif c == ')':
                depth -= 1

        return answer

s = Solution()
print(s.maxDepth('()()()()'))
print(s.maxDepth('()()(idfhgoiaudhgoishiudg())()'))
print(s.maxDepth('jdigdjbgueg236482'))
print(s.maxDepth('()((()()(())))'))
