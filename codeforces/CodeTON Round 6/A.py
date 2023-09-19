def solve():
    n, k, x = map(int, input().split())

    if x < k - 1:
        print(-1)
        return

    arr = []
    for i in range(k):
        arr.append(i)

    if len(arr) > n:
        print(-1)
        return

    for i in range(n - len(arr)):
        arr.append(x if x != k else x - 1)

    print(sum(arr))


t = int(input())

for _ in range(t):
    solve()
