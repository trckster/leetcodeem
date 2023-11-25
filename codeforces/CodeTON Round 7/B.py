def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a % b)


def solve():
    n = int(input())
    s = list(input())

    ans = n - 1

    i = 0
    while i < n and s[i] == 'B':
        i += 1
        ans -= 1
    i = n - 1
    while i >= 0 and s[i] == 'A':
        i -= 1
        ans -= 1

    print(max(0, ans))


t = int(input())
for _ in range(t):
    solve()
