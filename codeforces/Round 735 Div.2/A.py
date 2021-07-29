t = int(input())

for i in range(t):
    n = input()
    a = list(map(int, input().split(' ')))
    j = 0
    m = -1
    while j < len(a) - 1:
        x = a[j] * a[j + 1]
        m = max(x, m)
        j += 1
    print(m)
