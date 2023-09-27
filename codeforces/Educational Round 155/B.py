def solve():
    n = int(input())
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))

    am = min(a)
    bm = min(b)

    ap = am * n + sum(b)
    bp = bm * n + sum(a)
    print(min(ap, bp))


t = int(input())

for _ in range(t):
    solve()
