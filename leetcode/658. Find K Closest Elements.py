from typing import List


class Solution:
    def getClosestIndex(self, arr: List[int], x: int) -> int:
        if arr[0] > x:
            return 0
        elif arr[-1] < x:
            return len(arr) - 1

        l = 0
        r = len(arr) - 1
        while True:
            i = (l + r) // 2

            if x == arr[i]:
                return i

            if arr[i] < x < arr[i + 1]:
                leftDistance = x - arr[i]
                rightDistance = arr[i + 1] - x

                if leftDistance > rightDistance:
                    return i + 1
                else:
                    return i

            if x < arr[i]:
                r = i - 1
            else:
                l = i + 1

    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        result = []

        # Binary search the closest one
        l = r = self.getClosestIndex(arr, x)

        result.append(arr[l])
        k -= 1

        while k > 0:
            if l - 1 < 0:
                result += arr[r + 1:r + 1 + k]
                break
            elif r + 1 >= len(arr):
                result += arr[l - k:l]
                break

            leftDistance = x - arr[l - 1]
            rightDistance = arr[r + 1] - x

            if leftDistance > rightDistance:
                r += 1
                result.append(arr[r])
            else:
                l -= 1
                result.append(arr[l])

            k -= 1

        result.sort()

        return result


solution = Solution()
result = solution.findClosestElements([1, 2, 3, 4, 5, 8, 20, 200], 3, 6)
print(result)
