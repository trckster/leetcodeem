t = int(input())

for _ in range(t):
    a, b = map(int, input().split())

    if a == 1 or b == 1:
        print(1)
        print(a, b)
        continue

    print(2)
    print(a - 1, 1)
    print(a, b)
