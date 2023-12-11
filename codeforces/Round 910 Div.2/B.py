from math import ceil


def split(n, threshold):
    if threshold >= n:
        return 0, n

    places_count = ceil(n / threshold)
    minimal = n // places_count

    return places_count - 1, minimal


def solve():
    n = int(input())
    arr = list(map(int, input().split()))

    result = 0

    for i in range(len(arr) - 2, -1, -1):
        left = arr[i]
        right = arr[i + 1]
        n, minimum = split(left, right)
        result += n
        arr[i] = minimum

    print(result)


t = int(input())
for _ in range(t):
    solve()
