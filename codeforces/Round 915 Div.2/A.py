def yes():
    print('YES')


def no():
    print('NO')


def solve():
    a, b = map(int, input().split())
    print(max(a, b))


t = int(input())
for _ in range(t):
    solve()
