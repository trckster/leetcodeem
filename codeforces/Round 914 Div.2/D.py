def yes():
    print('YES')


def no():
    print('NO')


def solve():
    n = int(input())
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))
    tu = []
    for i in range(n):
        tu.append([
            a[i],
            b[i],
            i,
        ])
    tu.sort(key=lambda x: (x[1], x[0]))
    for tus in tu:
        i = tus[2]
        have = a[i]
        goal = b[i]
        if goal == have:
            continue
        if goal < have:
            no()
            return

        cured = False
        j = i - 1
        while j >= 0:
            cure = a[j]
            if cure > goal:
                break
            elif cure == goal:
                for ii in range(j + 1, i + 1):
                    a[ii] = cure
                cured = True
                break
            elif b[j] < goal:
                break
            j -= 1

        j = i + 1
        while j < n:
            cure = a[j]
            if cure > goal:
                break
            elif cure == goal:
                for ii in range(i, j):
                    a[ii] = cure
                cured = True
                break
            elif b[j] < goal:
                break
            j += 1

        if not cured:
            no()
            return
    yes()


t = int(input())
for _ in range(t):
    solve()
