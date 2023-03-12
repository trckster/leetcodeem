t = int(input())

for _ in range(t):
    n, x, p = map(int, input().split())

    answer = False

    for power in range(1, min(n * 2, p + 1)):
        skip = power * (power + 1) // 2
        res = (x + skip) % n

        if res == 0:
            answer = True
            break

    print('YES' if answer else 'NO')
