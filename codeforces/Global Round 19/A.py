t = int(input())

for i in range(t):
    n = int(input())
    l = list(map(int, input().split()))
    if all(l[i] <= l[i+1] for i in range(len(l) - 1)):
        print('NO')
    else:
        print('YES')
