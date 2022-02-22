def print_arr(arr):
    for i in arr:
        print(i, end=' ')
    print()

t = int(input())

for task in range(t):
    n = int(input())

    arr = []
    for i in range(n, 0, -1):
        arr.append(i)


    print_arr(arr)

    i = 0
    while i < n - 1:
        arr[-i - 1], arr[-i - 2] = arr[-i - 2], arr[-i - 1]
        i += 1
        print_arr(arr)
