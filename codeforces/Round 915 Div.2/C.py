def yes():
    print('YES')


def no():
    print('NO')


def solve():
    n = int(input())
    s = list(input())

    best = [s[-1]]
    best_indexes = [len(s) - 1]
    i = n - 2
    while i >= 0:
        if s[i] >= best[-1]:
            best.append(s[i])
            best_indexes.append(i)
        i -= 1
    best_indexes = best_indexes[::-1]
    i = 0
    for index in best_indexes:
        s[index] = best[i]
        i += 1

    good = True
    i = 1
    while i < len(s):
        if s[i - 1] > s[i]:
            good = False
            break
        i += 1

    i = len(best) - 2
    highest_count = 1
    while i >= 0:
        if best[i] == best[i + 1]:
            highest_count += 1
        else:
            break
        i -= 1

    if good:
        print(len(best_indexes) - highest_count)
    else:
        print(-1)


t = int(input())
for _ in range(t):
    solve()
