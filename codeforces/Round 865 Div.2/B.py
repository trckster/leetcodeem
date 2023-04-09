t = int(input())

for _ in range(t):
    n = int(input())

    maxima = n * 2

    x = 1

    for i in range(0, n, 2):
        print(maxima - x, x, end=' ')
        x += 2

    print()

    x = 2
    for i in range(0, n - 2, 2):
        print(x, maxima - x, end=' ')
        x += 2

    print(x, maxima)
