def get_min_or(arr):
    result = 0
    for i in arr:
        result = result | i
    return result

t = int(input())

for i in range(t):
    n = int(input())
    arr = list(map(int, input().split()))
    print(get_min_or(arr))