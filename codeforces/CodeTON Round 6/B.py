def xor(arr):
    res = 0
    for i in arr:
        res ^= i
    return res


def arr_or(arr):
    res = 0
    for i in arr:
        res |= i
    return res


def solve():
    n, m = map(int, input().split())
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))

    a_xor = xor(a)
    b_or = arr_or(b)

    if n % 2:
        print(a_xor, a_xor | b_or)
    else:
        print(a_xor & ~b_or, a_xor)


t = int(input())

for _ in range(t):
    solve()
