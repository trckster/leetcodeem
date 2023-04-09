def process(a):
    if len(a) % 2 == 1:
        return True

    i = 1
    while i < len(a) - 1:
        diff = arr[i] - arr[0]
        arr[i + 1] -= diff
        arr[i] = arr[0]
        i += 1
    return arr[-1] >= arr[-2]


t = int(input())

for _ in range(t):
    n = int(input())
    arr = list(map(int, input().split()))

    print('YES' if process(arr) else 'NO')
