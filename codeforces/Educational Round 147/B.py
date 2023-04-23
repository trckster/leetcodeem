def solve():
    n = int(input())
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))

    left = None
    for i in range(n):
        if a[i] != b[i]:
            left = i
            break

    right = None
    for i in range(n - 1, -1, -1):
        if a[i] != b[i]:
            right = i
            break

    while left >= 1 and b[left - 1] <= b[left]:
        left -= 1
    while right + 1 < n and b[right + 1] >= b[right]:
        right += 1

    print(left + 1, right + 1)


t = int(input())

for _ in range(t):
    solve()
