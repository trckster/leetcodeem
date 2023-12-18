def yes():
    print('YES')


def no():
    print('NO')


def solve():
    ab = input()

    i = 1
    while i < len(ab):
        a = ab[:i]
        b = ab[i:]

        if a[0] == '0' or b[0] == '0':
            i += 1
            continue

        aa = int(a)
        bb = int(b)

        if aa < bb:
            print(aa, bb)
            return
        i += 1
    print(-1)


t = int(input())
for _ in range(t):
    solve()
