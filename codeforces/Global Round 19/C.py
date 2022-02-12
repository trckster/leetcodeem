def get_rearrangements(arr):
    if all(i == 1 for i in arr) or len(arr) == 1 and arr[0] % 2:
        return -1

    answer = 0
    for num in arr:
        answer += (num + num % 2) // 2

    return answer


t = int(input())

for i in range(t):
    n = int(input())
    arr = list(map(int, input().split()))
    print(get_rearrangements(arr[1:-1]))
