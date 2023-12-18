def yes():
    print('YES')


def no():
    print('NO')


def solve():
    s = input()

    z = s.count('0')
    o = s.count('1')

    if z == o:
        print(0)
        return

    best = '1' if z < o else '0'

    first_len = len(s) - (max(z, o) - min(z, o))
    answer = max(z, o) - min(z, o)

    s = s[:first_len]
    best_count = s.count(best)

    missing = best_count - min(z, o)

    i = len(s) - 1
    while missing > 0 and i >= 0:
        answer += 1
        if s[i] == best:
            missing -= 1
        i -= 1

    print(answer)


t = int(input())
for _ in range(t):
    solve()
