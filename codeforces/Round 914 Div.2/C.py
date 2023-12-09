def yes():
    print('YES')


def no():
    print('NO')


def binary_search(arr, low, high, x):
    if high >= low:
        mid = (high + low) // 2

        if arr[mid] == x:
            return mid
        elif arr[mid] > x:
            return binary_search(arr, low, mid - 1, x)
        else:
            return binary_search(arr, mid + 1, high, x)
    else:
        return low


def get_min_diff(arr):
    diffs = []
    for i in range(1, len(arr)):
        diffs.append(abs(arr[i] - arr[i - 1]))
    return min(diffs)


def solve():
    n, k = map(int, input().split())
    arr = list(map(int, input().split()))

    if k >= 3:
        print(0)
    elif k == 2:
        arr.sort()
        first_diffs = set()
        for i in range(n):
            for j in range(i + 1, n):
                first_diffs.add(abs(arr[i] - arr[j]))

        default_min = min(arr)
        default_diffs_min = get_min_diff(arr)
        winner = min(default_min, default_diffs_min)

        for new in first_diffs:
            pos = binary_search(arr, 0, n - 1, new)
            possible_minimums = [new, abs(arr[pos] - new)]
            if pos + 1 < n:
                possible_minimums.append(abs(arr[pos + 1] - new))
            if pos - 1 >= 0:
                possible_minimums.append(abs(arr[pos - 1] - new))
            winner = min(winner, min(possible_minimums))
        print(winner)

    else:
        arr.sort()
        print(min(min(arr), get_min_diff(arr)))


t = int(input())
for _ in range(t):
    solve()
