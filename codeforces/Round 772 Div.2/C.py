def solve(arr):
    if arr[-1] < arr[-2]:
        print(-1)
        return

    if arr[-1] < 0:
        if all(arr[i] <= arr[i+1] for i in range(len(arr) - 1)):
            print(0)
        else:
            print(-1)
        return 

    k = arr[-2] - arr[-1]
    y = len(arr) - 1
    z = len(arr)

    print(len(arr) - 2)

    for i in range(1, len(arr) - 1):
        print(i, y, z)


t = int(input())

for i in range(t):
    n = int(input())
    arr = list(map(int, input().split()))
    solve(arr)
