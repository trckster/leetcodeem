MOD = 998244353


def fac(n):
    res = 1
    while n:
        res *= n
        res %= MOD
        n -= 1
    return res


def solve():
    s = input()
    arr = [1]
    curr = s[0]

    for i in range(1, len(s)):
        c = s[i]
        if curr == c:
            arr[-1] += 1
        else:
            arr.append(1)
            curr = c

    changes = 0
    for length in arr:
        changes += length - 1

    ways = 1
    for length in arr:
        if length != 1:
            ways *= length
            ways %= MOD

    ways *= fac(changes)
    ways %= MOD

    print(changes, ways)


t = int(input())

for _ in range(t):
    solve()
