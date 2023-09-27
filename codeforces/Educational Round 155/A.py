def solve():
    n = int(input())

    p = list(map(int, input().split()))
    answer = p[0]
    for _ in range(n - 1):
        s, e = list(map(int, input().split()))
        if s >= p[0] and e >= p[1]:
            answer = -1
    print(answer)


t = int(input())

for _ in range(t):
    solve()
