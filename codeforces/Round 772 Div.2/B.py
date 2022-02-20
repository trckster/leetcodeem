def solve(arr):
    changes_count = 0

    for i in range(1, len(arr) - 1):
        if arr[i - 1] < arr[i] and arr[i] > arr[i + 1]:
            changes_count += 1

            if i + 2 >= len(arr):
                arr[i + 1] = arr[i]
            else:
                arr[i + 1] = max(arr[i], arr[i + 2])
    print(changes_count)
    for element in arr:
        print(element, end=' ')
    print()


t = int(input())

for i in range(t):
    n = int(input())
    arr = list(map(int, input().split()))
    solve(arr)
