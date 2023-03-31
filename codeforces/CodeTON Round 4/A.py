def process(arr):
    for i in range(len(arr)):
        if arr[i] <= i + 1:
            return True
    return False


t = int(input())

for _ in range(t):
    n = int(input())
    arr = list(map(int, input().split()))

    print('yes' if process(arr) else 'no')
