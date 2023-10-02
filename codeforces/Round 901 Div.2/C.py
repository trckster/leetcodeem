powers_of_two = set([2 ** i for i in range(0, 50)])


def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a % b)


def solve():
    n, m = map(int, input().split())

    if n >= m:
        n %= m

    gd = gcd(m, n)

    if m // gd not in powers_of_two:
        print(-1)
        return

    answer = 0
    while n != 0:
        answer += n
        n *= 2
        n %= m

    print(answer)


t = int(input())
for _ in range(t):
    solve()
