def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a % b)


def solve():
    n = int(input())
    arr = list(map(int, input().split()))
    if arr[0] != 1:
        print('no')
        return

    print('yes')


t = int(input())
for _ in range(t):
    solve()
