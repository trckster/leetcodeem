def yes():
    print('YES')


def no():
    print('NO')


def save_positions(arr):
    backward = {}

    for i in range(len(arr)):
        if arr[i] in backward:
            backward[arr[i]].add(i)
        else:
            backward[arr[i]] = set()
            backward[arr[i]].add(i)

    return backward


def solve():
    n = int(input())
    arr = list(map(int, input().split()))
    position = save_positions(arr)
    arr.sort()
    answer = []
    current_sum = arr[0]
    last_set = 0
    for i in range(n - 1):
        if current_sum >= arr[i + 1]:
            current_sum += arr[i + 1]
        else:
            while last_set <= i:
                answer.append(i)
                last_set += 1
            current_sum += arr[i + 1]
    while last_set < n:
        answer.append(n - 1)
        last_set += 1

    real_answer = [None for i in range(n)]
    for i in range(n):
        curr_number = arr[i]
        curr_number_pos = position[curr_number].pop()
        real_answer[curr_number_pos] = str(answer[i])
    print(' '.join(real_answer))


t = int(input())
for _ in range(t):
    solve()
